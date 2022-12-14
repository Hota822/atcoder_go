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
	m := readInt()
	// s := read()

	// dp := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	dp[i] = make([]int, n)
	// }
	// mp := map[int][][]int{}

	xy := make([][]int, m)
	for i := 0; i < m; i++ {
		xy[i] = readSli(2)
	}

	sy := make([]int, n)

	for i := 0; i < 2; i++ {
		for _, v := range xy {

			x := v[0]
			y := v[1]

			// max_path = sy[x] + 1
			max_path := sy[x-1] + 1
			current_max_of_y := sy[y-1]
			// dp[x-1][y-1] = max_path
			if current_max_of_y < max_path {
				sy[y-1] = max_path
			}
			// p(x, y, sy, dp)
		}
	}

	// p(sy)
	ans := 0
	for _, v := range sy {
		if ans < v {
			ans = v
		}
	}

	return ans
}

// 4 5
// 1 2
// 1 3
// 3 2
// 2 4
// 3 4

// before
//    1 2 3 4 y sy
// 1  0 1 1 0
// 2  0 0 0 0
// 3  0 2 0 0
// 4  0 0 0 0
// x
//
// sy 0 1 1 0

// 5 8
// 5 3
// 2 3
// 2 4
// 5 2
// 5 1
// 1 4
// 4 3
// 1 3

// before
//    1 2 3 4 5 y  sx
// 1  0 0 2 2 0     2
// 2  0 0 2 2 0     2
// 3  0 0 0 0 0     0
// 4  0 0 3 0 0     3
// 5  1 1 1 0 0     1
// x
//
// sy 1 1 3 2 0

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
		fmt.Println(v)
	}
}
