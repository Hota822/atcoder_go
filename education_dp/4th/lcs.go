package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
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
)

var sc = bufio.NewScanner(os.Stdin)

// var dp [][]int
// var sli []int
// var memo [][]int

func run() {
	// n := readInt()
	s_sli := strings.Split(read(), "")
	t_sli := strings.Split(read(), "")

	dp := make([][]int, len(s_sli)+1)
	s_len := len(s_sli)
	t_len := len(t_sli)
	dp[0] = make([]int, t_len+1)
	// chr_count := make(map[string]int)

	for i := 1; i <= s_len; i++ {
		dp[i] = make([]int, t_len+1)
		s := s_sli[i-1]
		dp[i] = make([]int, t_len+1)
		// count := chr_count[s]
		// chr_count[s]++

		for j := 1; j <= t_len; j++ {
			t := t_sli[j-1]
			if s == t {

				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	len := dp[s_len][t_len]
	ans := make([]string, len)
	for i, j := s_len, t_len; len > 0; {
		if s_sli[i-1] == t_sli[j-1] {
			ans[len-1] = s_sli[i-1]
			i--
			j--
			len--
		} else if dp[i-1][j] == dp[i][j]-1 {
			j--
		} else {
			i--
		}
	}

	for _, v := range ans {
		fmt.Print(v)
	}
	fmt.Println()
}

//   a b r a c a d a b r a
// a 1 1 1 1 1 1 1 1 1 1 1
// v
// a 1 1 1 2 2 2 2 2 2 2 2
// d 1 1 1 2 2 2 3 3 3 3 3
// a 1 1 1 2 2 3 3 3 3 3 3
// k 1 1 1 2 2 3 3 3 3 3 3
// e 1 1 1 2 2 3 3 3 3 3 3
// d 1 1 1 2 2 3 4 4 4 4 4
// a 1 1 1 2 2 3 4 5 5 5 5
// v 1 1 1 2 2 3 4 5 5 5 5
// r 1 1 2 2 2 3 4 5 5 6 6
// a 1 1 2 3 3 3 4 5 6 6 7
//   a b r a c a d a b r a

// aaadara

//   a x y b
// a 1 1 1 1
// b 1 1 1 2
// y 1 1 2 2
// x 1 2 2 2
// b 1 2 2 3

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

// =======================main========================
func main() {
	buf := make([]byte, initial_buf)
	sc.Buffer(buf, max_bufSize)
	sc.Split(bufio.ScanWords)
	run()
	// result := run()
	// print(result)
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

func s(args ...interface{}) []string {
	var s_sli []string
	for _, v := range args {
		if s, ok := v.(string); ok {
			s_sli = append(s_sli, s)
		}
	}
	return s_sli
}
