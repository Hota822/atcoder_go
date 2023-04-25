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
	max_int64 = 9223372036854775807
	// prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

var dp [][]int
var sli []int
var memo [][]bool
var cum []int

func run() interface{} {
	n := readInt()

	sli = readSli(n)
	dp = make([][]int, n)
	memo = make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		memo[i] = make([]bool, n)
	}

	cum = make([]int, n+1)
	for i := 1; i <= n; i++ {
		cum[i] = cum[i-1] + sli[i-1]
	}

	ans := calculate(0, n-1)
	return ans
}

func calculate(l, r int) int {
	// p(dp);
	if memo[l][r] {
		return dp[l][r]
	}
	memo[l][r] = true

	if l == r {
		return 0
	}

	min := max_int64
	for i := l; i < r; i++ {
		// 	pl("l, i: ", l, i, "   i+1, r: ", i+1, r, "cum: ", cum[r+1]-cum[i])
		min = Min(min, calculate(l, i)+calculate(i+1, r))
	}

	dp[l][r] = min + cum[r+1] - cum[l]

	return dp[l][r]
}

// 4
// 10 20 30 40

// 10 + [20 30 40]
//   > 20 + [30 40]
//     > 30 + 40
//   > [20 30] + 40
//     > 30 + 40
// [10 20] + [30 40]
//   > 10 + 20
// [10 20 30] + 40
//   > 10 + [20 30]
//   > [10 20] + 30

// [10 0 0 0]
// [0 20 0 0]
// [0 0 30 70]
// [0 0 0 40]

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
