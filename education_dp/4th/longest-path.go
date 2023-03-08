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

// var dp [][]int
// var sli []int
// var memo [][]int

func run() interface{} {
	n, m := readInt(), readInt()
	// s := read()

	sli := make([][]int, m)
	for i := 0; i < m; i++ {
		sli[i] = readSli(2)
	}

	dp := make([]int, n)

	a := initGraph(n, sli)
	a.TopologicalSort()
	for _, paths := range a.sorted {
		for _, path := range paths {
			from, to := path[0]-1, path[1]-1
			dp[to] = Max(dp[from]+1, dp[to])
		}
	}
	p(dp)
	ans := 0
	for _, i := range dp {
		ans = Max(ans, i)
	}
	return ans
}

// 99999

type Graph struct {
	vert   int
	paths  [][]int
	idx    []int
	sorted [][][]int
}

func initGraph(vert int, paths [][]int) *Graph {
	sli := make([]int, vert)
	for i := 0; i < vert; i++ {
		sli[i] = i + 1
	}
	sorted := make([][][]int, vert)
	grh := Graph{vert: vert, idx: sli, paths: paths, sorted: sorted}
	return &grh
}

func (grh *Graph) TopologicalSort() {
	// sli := make([]int, grh.vert)
	// copy(sli, grh.idx)
	for _, item := range grh.paths {
		from, to := item[0]-1, item[1]-1
		from_idx, to_idx := grh.idx[from]-1, grh.idx[to]-1
		grh.sorted[from_idx] = append(grh.sorted[from_idx], item)
		if from_idx > to_idx {
			// sli[from_idx], sli[to_idx] = sli[to_idx], sli[from_idx]
			grh.sorted[from_idx], grh.sorted[to_idx] = grh.sorted[to_idx], grh.sorted[from_idx]
			grh.idx[from], grh.idx[to] = grh.idx[to], grh.idx[from]
		}
	}
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
