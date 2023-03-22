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
)

var sc = bufio.NewScanner(os.Stdin)

var dp [][][]float64

// var memo map[string]bool
var c []int
var n int

func run() interface{} {
	n = readInt()

	c := make([]int, 4)
	for i := 0; i < n; i++ {
		idx := readInt()
		c[idx]++
	}

	dp = make([][][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]float64, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]float64, n+1)
		}
	}
	dp[1][0][0] = float64(n)

	ans := calculate(0, c[1], c[2], c[3])
	return ans
}

func calculate(c0, c1, c2, c3 int) float64 {
	if dp[c1][c2][c3] != 0 {
		return dp[c1][c2][c3]
	}

	// 期待値 = 0がない場合：それ以下の期待値＋１
	//            ある場合：それ以下の期待値 + (残っている皿の数/全数)
	var e float64
	if c0 == 0 {
		e = 1
	} else {
		e = float64(n) / float64(n-c0)
	}

	// それ以外の期待値 = 各cが１以上の時に計算可能
	e1 := 0.0
	if c1 > 0 {
		e1 = float64(c1) / float64(n-c0) * calculate(c0+1, c1-1, c2, c3)
	}

	e2 := 0.0
	if c2 > 0 {
		e2 = float64(c2) / float64(n-c0) * calculate(c0, c1+1, c2-1, c3)
	}

	e3 := 0.0
	if c3 > 0 {
		e3 = float64(c3) / float64(n-c0) * calculate(c0, c1, c2+1, c3-1)
	}

	dp[c1][c2][c3] = e + e1 + e2 + e3
	return dp[c1][c2][c3]
}

// 2
// 1 2
// 0 1 1 0 (110)
// 1 0 1 0 (010)
// 2 1 0 0 (100)
// 0 2 0 0 (200)

// 1 2
// > 1 1 = cal(01) + 1 = 2 + 1 = 3
//   > 0 1 = 2
// > 0 2 = cal(01) + 2 = 2 + 2
//   > 0 1 = 2

// 3
// 1 2 3

// 0 1 1 1
// 1 0 1 1 (0 2 3) > c1 -1, c0 +1
// 0 2 0 1 (1 1 3) > c2 -1, c1 +1
// 0 1 2 0 (1 2 2) > c3 -1, c2 +1

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
			}
			continue
		}
		if dp, ok := v.([][]float64); ok {
			for _, v := range dp {
				fmt.Print(v)
			}
			continue
		}
		fmt.Print(v)
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
