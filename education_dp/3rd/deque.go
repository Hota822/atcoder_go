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
var sli []int
var xy map[bool]int
var dp [][]int
var flag [][]int

func run() interface{} {
	n := readInt()
	// s := read()

	sli = readSli(n)
	if n == 1 {
		return sli[0]
	}

	// left end and right end dp
	dp = make([][]int, n)
	flag = make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		flag[i] = make([]int, n)
	}

	// x1, y1 := calculate(1, n-1)
	// x2, y2 := calculate(0, n-2)

	// abs1 := Abs(x1 - y1)
	// abs2 := Abs(x2 - y2)
	// return Max(abs1, abs2)
	ans := calculate(0, n-1)
	return ans
}

// 4
// 10 80 90 30
// lr = (0,3)
// lx, ly = calc(1, 3)
//  = 10 + calc(1,3)
// calc(1,3) = (lx, ly) = 80 + x, calc(2,3) or 30 + x, calc(1,2)
// rx, ry = calc(0, 2)

func calculate(l, r int) int {
	// p([]int{l, r})
	if flag[l][r] == 1 {
		return dp[l][r]
	}

	if r == l {
		dp[l][r] = sli[l]
		return sli[l]
	}

	// p([]int{l_y, l_x})
	flag[l][r] = 1

	dp[l][r] = Max(sli[l]-calculate(l+1, r), sli[r]-calculate(l, r-1))
	return dp[l][r]
}

// 429715

// get l(4) > 29715
//   get l(2) > 9715
//     get l(9) > 715
//       get l(7) > 15
//         get 5 > get 1
//       get r(5) > 71
//          get 7 > get 1
//     get r(5) > 971
//       get l(9) >71 same
//       get 1(1) >97
//   get r(5) > 2971
//     get l(2) > 971 same
//     get r(1) > 297
//       get l(2) > 97 same
//       get r(7) > 29
//         get 9 > get 2
// get r(5) > 42971
//   get l(4) > 2971 same
//   get r(1) > 4297
//     get l(4) > 297 same
//     get r(7) > 429
//       get l(4) > 29 same
//       get r(9) > 42
//         get 4 > get 2

// (715)
// 7 + 1, 5 or 5 + 1, 7
// 8-5=3, 6-7=-1
//
// (971)
// 9 + 1, 7 or 1 + 7, 9
// 10-7=3, 8-9=-1
// 10, 7

// (9715)
// get l(9) > 715 or get r(5) > 971
//   get l > 9 + 715(1), 715(0) > 9 + 5, 8
//   get r > 5 + 971(1), 971(0) >

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
