package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
	// max_int64 = 9223372036854775807
	prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	n := readInt()

	sli := make([][]int, n)
	for i := 0; i < n; i++ {
		sli[i] = readSli(n)
	}

	dp := make([]int, 1<<n)
	dp[0] = 1
	for S := 0; S < 1<<n; S++ {
		if dp[S] == 0 {
			continue
		}

		// fmt.Print("S: ")
		// fmt.Print(S)
		// fmt.Print(", S(2): ")
		// fmt.Printf("%b", S)
		// fmt.Println()

		// 立っているビット数をカウント > マッチした人数
		// i = 0, 1, 2, 3, 4...
		i := bits.OnesCount(uint(S))
		for j := 0; j < n; j++ {
			// e.g. S=0100, j=1
			// S>>j = 0010, S>>j&1 = 0, S>>j&1 == false
			if S>>j&1 == 1 {
				if i == 1 {

					fmt.Print("S: ")
					fmt.Print(S)
					fmt.Print(", S(2): ")
					fmt.Printf("%b", S)
					fmt.Print(", j: ")
					fmt.Print(j)
					fmt.Print(", S>>j: ")
					fmt.Printf("%b", S>>j)
					fmt.Println()
					// ps([]string{"S", "S>>j", "S>>j&1"}, S, S>>j, S>>j&1)
				}
				continue
			}

			// if i == 0 {
			// 	p(1)
			// }
			if sli[i][j] == 1 {
				dp[S|1<<j] += dp[S]
				dp[S|1<<j] %= prime_number
			}
		}
	}

	ans := dp[1<<n-1]

	return ans
}

// S==0, into j loop, set

// 4
//i\j
// 0 1 0 0
// 0 0 0 1
// 1 0 0 0
// 0 0 1 0

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
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
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
		if dp, ok := v.([][]float64); ok {
			for _, v := range dp {
				fmt.Println(v)
			}
			continue
		}
		fmt.Println(v)
	}
}

func ps(header []string, arg ...interface{}) {
	_, _, l, _ := runtime.Caller(1)
	s := strconv.Itoa(l)
	fmt.Println("dumped at line: " + s + ", value: ")
	for idx, v := range arg {
		fmt.Print(header[idx])
		fmt.Print(": ")
		fmt.Print(v)
		fmt.Print(", ")
	}
	fmt.Println()
}
