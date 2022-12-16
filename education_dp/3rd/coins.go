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
	n := readInt()
	// s := read()
	p_sli := make([]float64, n+1)

	// sli := make([][]int, n)
	dp := make([][]float64, n+1)
	dp[0] = make([]float64, n+1)
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		p_sli[i] = readFloat()
	}

	// i枚目のコインまでで
	for i := 1; i <= n; i++ {
		dp[i] = make([]float64, n+1)
		// j枚表が出る確率を記録
		for j := 0; j <= n; j++ {
			if j == 0 {
				// 全て裏
				dp[i][0] = dp[i-1][0] * (1 - p_sli[i])
			} else {
				// 今回が表 > 前回まで j -1の数出ていて、今回表
				t := dp[i-1][j-1] * p_sli[i]
				// 今回が裏 > 前回まで j出ていて、今回裏
				f := dp[i-1][j] * (1 - p_sli[i])
				dp[i][j] = t + f
			}

		}
	}

	// ans := p_sli
	ans := 0.0
	require := n/2 + 1
	for i := n; i >= require; i-- {
		ans += dp[n][i]
	}

	return ans
}

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

func readFloat() float64 {
	sc.Scan()
	ret, _ := strconv.ParseFloat(sc.Text(), 64)
	return ret
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
