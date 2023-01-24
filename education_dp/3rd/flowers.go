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

var dp [][]int
var sli []int
var memo [][]int

func run() interface{} {
	n := readInt()
	// s := read()

	h_sli := readSli(n)
	a_sli := readSli(n)

	type Stats = struct {
		max_height int
		sum_value  int
		// diff       int
	}

	// dp := map[int]Stats{}
	// dp := make([]Stats, n+1)
	dp := newSegmentTree(n + 1)

	// max_h := 0
	for i := 0; i < 0; i++ {
		// sum := 0
		// before := 0

		// for j := i; j < n; j++ {
		// 	height := h_sli[j]
		// 	value := a_sli[j]

		// }

	}

	for i := 0; i < n; i++ {
		// p(i, dp[i].sum_value)
		h, a := h_sli[i], a_sli[i]
		dp.update(i, h)
	}
}

type SegmentTree struct {
	data        []int
	element_num int
}

func newSegmentTree(n int) SegmentTree {
	n *= 2
	st := SegmentTree{data: make([]int, n), element_num: n}
	return st
}

func getMax(x, y int, tree SegmentTree) int {
	var get func(x, y, l, r, k int) int
	get = func(x, y, l, r, k int) int {
		if x <= l && r <= y { // 範囲内
			return tree.data[k]
		} else { // 一部被っているとき。範囲外は考えない
			max_l := get(x, y, l, (l+r)/2, k*2+1)
			max_r := get(x, y, (l+r)/2, r, k*2+2)
			return Max(max_l, max_r)
		}
	}
	return get(x, y, 0, tree.element_num, 0)
}

func (tree SegmentTree) update(idx, val int) {
	idx += tree.element_num - 1 // i番目の葉は i + n-1に存在する
	tree.data[idx] = val        // 葉のデータ更新
	for i := idx; i > 0; {
		i = (i - 1) / 2                                        // 親のidx取得
		tree.data[i] = Max(tree.data[i*2+1], tree.data[i*2+2]) // 子を確認し、更新
	}
}

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	var n int
// 	fmt.Fscan(in, &n)
// 	hs := make([]int, n)
// 	as := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		fmt.Fscan(in, &hs[i])
// 	}
// 	for i := 0; i < n; i++ {
// 		fmt.Fscan(in, &as[i])
// 	}

// 	dp := newBIT(n + 1)
// 	for i := 0; i < n; i++ {
// 		h, a := hs[i], as[i]
// 		dp.update(h, dp.get(h-1)+a)
// 	}
// 	fmt.Println(dp.get(n))
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// type BIT struct {
// 	n   int
// 	bit []int
// }

// func newBIT(n int) BIT {
// 	b := BIT{n: n}

// 	n2 := 1
// 	for n2 <= n {
// 		n2 *= 2
// 	}

// 	b.bit = make([]int, n2)
// 	return b
// }

// func (b *BIT) get(i int) int {
// 	s := 0
// 	for i > 0 {
// 		s = max(s, b.bit[i])
// 		i -= i & -i
// 	}
// 	return s
// }

// func (b *BIT) update(i, x int) {
// 	for i <= b.n {
// 		b.bit[i] = max(b.bit[i], x)
// 		i += i & -i
// 	}
// }

//
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
