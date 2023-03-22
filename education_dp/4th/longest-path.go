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

type Node struct {
	self int              // 自身の番号
	next map[int]struct{} // 隣接している、矢印の先の頂点の番号
	deg  int              // 入次数
}

func run() interface{} {
	n, m := readInt(), readInt()

	// グラフの初期化
	graph := make(map[int]*Node)
	for i := 1; i <= n; i++ {
		graph[i] = &Node{next: map[int]struct{}{}, deg: 0, self: i}
	}
	// 経路の読込
	for i := 0; i < m; i++ {
		from, to := readInt(), readInt()
		graph[from].next[to] = struct{}{}
		graph[to].deg++
	}

	// トポロジカルソート実行----------------------
	// 既に処理可能なノードをキューに入れる
	var queue []*Node
	for _, node := range graph {
		if node.deg == 0 {
			queue = append(queue, node)
		}
	}

	// キューの先頭から処理していく
	for i := 0; i < len(queue); i++ {
		node := queue[i]
		// グラフから頂点を削除する
		delete(graph, node.self)

		for idx := range node.next {
			// 矢印の先の入次数を減らす
			(graph[idx].deg)--
			// 入次数が0のとき、キューに追加する
			if graph[idx].deg == 0 {
				queue = append(queue, graph[idx])
			}
		}
	}
	// トポロジカルソート終了----------------------

	// 最大値を保持し、計算する
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		node := queue[i]
		for j := range node.next {
			dp[j-1] = Max(dp[j-1], dp[node.self-1]+1)
		}
	}

	// 全ての中での最大値を取得する
	ans := 0
	for i := 0; i < n; i++ {
		ans = Max(ans, dp[i])
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
