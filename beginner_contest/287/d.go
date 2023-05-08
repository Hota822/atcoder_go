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
	// max_int64 = 9223372036854775807
	// prime_number = 1000_000_007
	forward  = 0
	backward = 1
)

var sc = bufio.NewScanner(os.Stdin)

// var dp [][]int
// var sli []int
// var memo [][]int

func run() interface{} {
	// n := readInt()
	s, t := read(), read()
	t_len := len(t)
	// len, forward or backward, match or not
	dp := make([][]bool, t_len+1)
	dp[0] = make([]bool, 2)
	dp[0][forward] = match(s[0], t[0])
	dp[0][backward] = match(s[]

	// s_sli := strings.Split(s, "")
	ans := make([]bool, 0)
	for x := 1; x <= t_len+1; x++ {
		// p(x)
		dp[x] = make([]bool, 2)
		forward_str := s[:x]
		backward_str := s[x:]
		// pl(forward_str, backward_str)

		for x < len(forward_str) {
			dp[x][forward] = dp[x-1][forward] && match(forward_str[x], t[x])
		}
		for x < len(backward_str) {
			dp[x][backward] = dp[x-1][backward] && match(backward_str[len(forward_str)+x], t[x])
		}
	}
	for i := 0; i < t_len; i++ {
		ans = append(ans, dp[i][forward] && dp[t_len-i][backward])
	}

	// ans := sli
	return ans
}

func match(a, b byte) bool {
	if string(a) == "?" || string(b) == "?" {
		return true
	}

	return a == b
}

// ========================read
func read() string {
	sc.Scan()
	return sc.Text()
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
	if sli, ok := ans.([]bool); ok {
		for _, v := range sli {
			if v {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
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
		// pointer
		// if dp, ok := v.([]*Rope); ok {
		// 	fmt.Print("[ ")
		// 	for i, v := range dp {
		// 		if i == 0 {
		// 			continue
		// 		}
		// 		fmt.Print(*v)
		// 		fmt.Print(" ")
		// 	}
		// 	fmt.Println("]")
		// 	continue
		// }
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
