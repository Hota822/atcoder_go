package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	// "strings"
	// "math"
	// "reflect"
	// "sort"
)

const (
	max_bufSize = 1_000_000_000 // default: 65536
	initial_buf = 10000
	// max_int32 = 2147483647
	// prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	n, k := readInt(), readInt()
	// s := read()

	sli := readSli(n)

	dp := make([][]bool, k+1)
	// dp[0] = make([]bool, n)

	for i := 0; i <= k; i++ {
		dp[i] = make([]bool, n+1)
		wl := false
		for j := 0; j < n; j++ {
			remain := i - sli[j]
			if remain < 0 {
				dp[i][j] = false
				continue
			}
			result := !dp[remain][n]
			// if i == 3 {
			// 	p([]int{remain, j}, []bool{result, wl})
			// }
			dp[i][j] = result || wl
			wl = dp[i][j]
		}
		dp[i][n] = wl
	}
	// p(dp)
	// dpf := make([][]string, k+1)
	// for i := 0; i <= k; i++ {
	// 	dpf[i] = make([]string, n+1)
	// 	for j := 0; j <= n; j++ {
	// 		if dp[i][j] {
	// 			dpf[i][j] = "w"
	// 		} else {
	// 			dpf[i][j] = "l"
	// 		}
	// 	}
	// }
	// p(dpf)
	ans := dp[k][n]
	return ans
}

//    1 2 3 r
// 0 [l l l l]
// 1 [w l l w]
// 2 [l w l w]
// 3 [w w w w]
// 4 [l l w w]
// 5 [w w w w]
// 6 [l w w w]
// 7 [w w w w]
// 8 [l l l l]
// 9 [w w w w]
//10 [l w w w]
//11 [w w w w]
//12 [l l l l]
//13 [w w w w]
//14 [l w w w]
//15 [w w w w]
//16 [l l l l]
//17 [w w w w]
//18 [l w w w]
//19 [w w w w]
//20 [l l l l]

// 2 - 2 = 0 -> win  ... A ()
// 2 -3 < 0 -> look left
// 3 -2 = 1 -> look 1 -> win
// 3 -3 = 0 -> win ... A
// 4 -2 = 2 -> look 2 -> w -> l
// 4 -3 = 1 -> look 2 -> w -> r is w (even 1 w is exist)
// 5
// ========================read
// func read() string {
// 	sc.Scan()
//     return sc.Text()
// }

// func readSli(n int) []string {
func readSli(n int) []int {
	// sli := make([]string, n)
	sli := make([]int, n)
	for i := 0; i < n; i++ {
		// sli[i] = read()
		sli[i] = readInt()
	}
	return sli
}

func readInt() int {
	sc.Scan()
	ret, _ := strconv.Atoi(sc.Text())
	return ret
}

// =======================main========================
func main() {
	buf := make([]byte, initial_buf)
	sc.Buffer(buf, max_bufSize)
	sc.Split(bufio.ScanWords)
	result := run()
	print(result)
}

func print(ans interface{}) {
	if v, ok := ans.(bool); ok {
		if v {
			fmt.Println("First")
		} else {
			fmt.Println("Second")
		}
		return
	}
	if sli, ok := ans.([]int); ok {
		for _, v := range sli {
			fmt.Println(v)
		}
		return
	}
	if sli, ok := ans.([]int64); ok {
		for _, v := range sli {
			fmt.Println(v)
		}
		return
	}
	if sli, ok := ans.([]string); ok {
		for _, v := range sli {
			fmt.Println(v)
		}
		return
	}
	fmt.Println(ans)
}

// =============================math
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func p(arg ...interface{}) {
	_, _, l, _ := runtime.Caller(1)
	s := strconv.Itoa(l)
	fmt.Println("dumped at line: " + s + ", value: ")
	for _, v := range arg {
		if dp, ok := v.([][]int); ok {
			for _, v := range dp {
				fmt.Println(v)
			}
			continue
		}
		if dp, ok := v.([][]string); ok {
			for _, v := range dp {
				fmt.Println(v)
			}
			continue
		}
		fmt.Println(v)
	}
}
