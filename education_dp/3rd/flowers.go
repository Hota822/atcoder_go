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

var dp SegmentTree
var sli []int
var memo [][]int

func run() interface{} {
	n := readInt()

	h_sli := readSli(n)
	a_sli := readSli(n)

	dp := newSegmentTree(n + 1)

	for i := 0; i < n; i++ {
		// p(i, dp[i].sum_value)
		h, a := h_sli[i], a_sli[i]
		// p(h)
		dp.update(h, dp.getMax(h-1)+a)
		// p(dp)

	}
	return dp.getMax(n)
}

type SegmentTree struct {
	data        []int
	element_num int
}

func newSegmentTree(n int) SegmentTree {
	st := SegmentTree{element_num: n}
	st.data = make([]int, n+1)
	// p(len(st.data), n)
	return st
}

func (tree SegmentTree) getMax(idx int) int {
	max := 0
	for i := idx; i > 0; {
		max = Max(max, tree.data[i])
		// i -= i & -i でセグメント木のNodeにアクセス
		i -= i & -i
	}
	return max
}

func (tree SegmentTree) update(idx, val int) {
	for i := idx; i <= tree.element_num; {
		tree.data[i] = Max(val, tree.data[i])
		i += i & -i
	}
}

// h[4 2 5 8 3 6 1 7 9]
// a[6 8 8 4 6 3 5 7 5]
// h=4
//  0 1 2 3 4 5 6 7 8 9 0
// [0 0 0 0 6 0 0 0 0 0 0] // dp[h]を更新
// [0 0 0 0 6 0 0 0 6 0 0] // ?

// h=2
//  0 1 2 3 4 5 6 7 8 9 0
// [0 0 8 0 6 0 0 0 6 0 0] // dp[h]を更新
// [0 0 8 0 8 0 0 0 6 0 0] // ? dp[4]
// [0 0 8 0 8 0 0 0 8 0 0] // ? dp[8]

// h=5
//  0 1 2 3 4 5  6  7 8 9 0
// [0 0 8 0 8 16 0  0 8 0 0] // dp[h]を更新
// [0 0 8 0 8 16 16 0 8 0 0] // ? dp[6]
// [0 0 8 0 8 16 16 0 16 0 0] // ? dp[8]

// h=8
//  0 1 2 3 4 5  6  7 8  9 0
// [0 0 8 0 8 16 16 0 20 0 0] // dp[h]を更新

// h=3
//  0 1 2 3  4  5  6  7 8  9 0
// [0 0 8 14 8  16 16 0 20 0 0] // dp[h]を更新
// [0 0 8 14 14 16 16 0 20 0 0] // ? dp[4]
// [0 0 8 14 14 16 16 0 20 0 0] // ? dp[8]

// h=6
//  0 1 2 3  4  5  6  7 8  9 0
// [0 0 8 14 14 16 19 0 20 0 0] // dp[h]を更新
// [0 0 8 14 14 16 19 0 20 0 0] // ? dp[8]

// h=1
//  0 1 2 3  4  5  6  7 8  9 0
// [0 5 8 14 14 16 19 0 20 0 0] // dp[h]を更新
// [0 5 8 14 14 16 19 0 20 0 0] // ? dp[2]
// [0 5 8 14 14 16 19 0 20 0 0] // ? dp[4]
// [0 5 8 14 14 16 19 0 20 0 0] // ? dp[8]

// h=7
//  0 1 2 3  4  5  6  7 8  9 0
// [0 5 8 14 14 16 19 26 20 0 0] // dp[h]を更新
// [0 5 8 14 14 16 19 26 26 0 0] // dp[8]

// h=7
//  0 1 2 3  4  5  6  7  8  9  0
// [0 5 8 14 14 16 19 26 26 31 0] // dp[h]を更新
// [0 5 8 14 14 16 19 26 26 31 31] ? dp[10]

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