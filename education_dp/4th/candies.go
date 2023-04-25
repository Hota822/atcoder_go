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
	max_int32   = 2147483647
	// max_int64 = 9223372036854775807
	prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

// var dp [][]int
// var sli []int
// var memo [][]int

func run() interface{} {
	n, k := readInt(), readInt()
	// s := read()

	sli := readSli(n)
	dp := make([][]int, n+1)
	dp[0] = make([]int, k+1)
	dp[0][0] = 1
	// dp := initSliceAs(k, 1)

	for i := 1; i <= n; i++ {
		dp[i] = make([]int, k+1)
		cum := make([]int, k+2)

		cum[0] = 0
		for j := 1; j <= k+1; j++ {
			cum[j] = (cum[j-1] + dp[i-1][j-1]) % prime_number
		}

		for j := 0; j <= k; j++ {
			dp[i][j] = (cum[j+1] - cum[Max(0, j-sli[i-1])] + prime_number) % prime_number
		}
	}

	ans := dp[n][k]

	return ans
}

// i番目の子供までで、ｊ個の飴を分け合う
// i=0, j=0, dp=1
// i=0, j=1, dp=0, ...

// i=1, j=0, dp=1 (0)
// i=1, j=1, dp=1 (1)
// i=1, j=2, dp=0
// i=1, j=3, dp=0
// i=1, j=4, dp=0

// i=2, j=0, dp=1
// i=2, j=1, dp=2 (0,1)(1,0)
//   > dp[i-1][j-1] +dp[i-1][j] = 2
//   > dp[i-1][0] + dp[j-1][1]
// i=2, j=2, dp=1 (1,1)(0,2)
//   > dp[i-1][j-2] + dp[i-1][j-1] + dp[i-1][j]
//   > dp[i-1][0] + dp[i-1][1] + dp[i-1][2] = 2
// i=2, j=3, dp=1 (1,2)
//   > dp[i-1][j-3] + dp[i-1][j-2] + dp[i-1][j-1] + dp[i-1][j] =1
//   >     ×             1               0             0  = 1
// i=2, j=4, dp=0
//   > dp[i-1][j]

// i=3, j=0, dp=1
// i=3, j=1, dp=3 (1,0,0)(0,1,0)(0,0,1)
// i=3, j=2, dp=3 (1,1,0)(1,0,1)(0,1,1)
// i=3, j=3, dp=6 (1,2,0)|(1,1,1)(0,2,1)|(1,0,2)(0,1,2)|(0,0,3)
// i=3, j=4, dp=5 (0,1,3)(0,2,2)(1,0,3)(1,1,2)(1,2,1)

// ========================read
// func read() string {
// 	sc.Scan()
//     return sc.Text()
// }

func initSliceAs(len, val int) []int {
	sli := make([]int, len)
	sli[0] = val
	for i := 1; i < len; i *= 2 {
		copy(sli[i:], sli[:i])
	}
	return sli
}

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

func readFloat() float64 {
	sc.Scan()
	ret, _ := strconv.ParseFloat(sc.Text(), 64)
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

func pr(arg ...interface{}) {
	for _, v := range arg {
		if dp, ok := v.([][]int); ok {
			for _, v := range dp {
				fmt.Print(v)
				fmt.Print(", ")
			}
			continue
		}
		if dp, ok := v.([][]float64); ok {
			for _, v := range dp {
				fmt.Print(v)
				fmt.Print(", ")
			}
			continue
		}
		fmt.Print(v)
		fmt.Print(", ")
	}
	fmt.Println()
}

func pl(arg ...interface{}) {
	fmt.Println()

	pr(arg...)
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

func s(args ...interface{}) []string {
	var s_sli []string
	for _, v := range args {
		if s, ok := v.(string); ok {
			s_sli = append(s_sli, s)
		}
	}
	return s_sli
}
