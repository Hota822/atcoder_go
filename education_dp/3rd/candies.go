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
	prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	n, k := readInt(), readInt()
	// s := read()

	sli := readSli(n)
	// var sum []int
	// var sli []int
	// var s int
	// for i:=0; i<n; i++ {
	//     v := readInt()
	//     sli[i] = v
	//     s += v
	//     sum[i] = s
	// }
	// if k > sum {
	//     return 0
	// }

	dp := make([][]int, n+1)
	dp[0] = make([]int, k+1)
	dp[0][0] = 1
	// dp[0] = initSliceAs(k+1, 1)
	sum := make([]int, k+2)

	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)
		sum[0] = 0 // banpei
		for j := 1; j <= k+1; j++ {
			sum[j] = (sum[j-1] + dp[i-1][j-1]) % prime_number
		}
		for j := 0; j <= k; j++ {
			dp[i][j] = (sum[j+1] - sum[Max(0, j-sli[i-1])] + prime_number) % prime_number
		}
	}

	ans := dp[n][k]
	return ans
}

// 3 4
// i番目の人、j個の飴
// i\j
//   0 1 2 3 4
// 0 1 0 0 0 0
// 1 1 1 0 0 0
// 2 1 2
// 3

// 2,0 > (0,0) =1
//  sum> [1]
// 2,1 > (1,0),(0,1)
//  sum> [1,2]
// 2,2 > (1,1),(0,2)
//  sum> [1,2,]

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

func initSliceAs(len, val int) []int {
	sli := make([]int, len)
	sli[0] = val
	for i := 1; i < len; i *= 2 {
		copy(sli[i:], sli[:i])
	}
	return sli
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
