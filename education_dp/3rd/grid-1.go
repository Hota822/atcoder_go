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
	// prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	h := readInt()
	w := readInt()
	// s := read()
	prime := 1000_000_007

	sli := make([][]string, h)
	for i := 0; i < h; i++ {
		sli[i] = strings.Split(read(), "")
	}

	dp := make([][]int, h+1)
	// dp[0] = make([]int, w)

	// lastSharp := 0
	dp[0] = make([]int, w+1)
	dp[0][1] = 1
	for i := 1; i <= h; i++ {
		dp[i] = make([]int, w+1)
		for j := 1; j <= w; j++ {
			if sli[i-1][j-1] == "." {
				dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % prime
			} else {
				dp[i][j] = 0
			}
		}
	}
	//   1 2 3 4
	// 1 1 1 1 0
	// 2 1 0 1 1
	// 3 1 1 2 3

	// 1列目：１
	//      ＃を保持
	// ２列目：
	//      左、上を見て和を取る
	//

	// sum

	// p(dp)
	ans := dp[h][w]
	return ans
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
		fmt.Println(v)
		if dp, ok := v.([][]int); ok {
			for _, v := range dp {
				fmt.Println(v)
			}
		}
	}
}
