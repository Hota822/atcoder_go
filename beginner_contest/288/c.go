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
	links []int
}

// map[root node]Tree
var tree map[int]*Node
var appear map[int]struct{}
var memo map[int]struct{}
var ans int

func run() interface{} {
	_, m := readInt(), readInt()
	// s := read()
	if m == 0 {
		return 0
	}

	tree = make(map[int]*Node)
	memo = make(map[int]struct{})
	for i := 0; i < m; i++ {
		a, b := readInt(), readInt()
		createNode(a, b)
		createNode(b, a)
		memo[a] = struct{}{}
		memo[b] = struct{}{}
	}

	// true if vertex have already appeared
	appear = make(map[int]struct{})
	appear[1] = struct{}{}
	ans = 0
	var node int
	for len(memo) != len(appear) {

		for node = range tree {
			p(node, len(tree))
			calculate(node, 0)
			break
		}
	}

	return ans
}

func calculate(current, parent int) {
	node := tree[current]
	delete(tree, current)
	children := node.getChildren(parent)
	// p(current)
	for _, v := range children {
		if _, ok := appear[v]; ok {
			// delete link
			ans++
		} else {
			appear[v] = struct{}{}
			calculate(v, current)
		}
	}
}

func createNode(parent, child int) {
	if node, ok := tree[parent]; ok {
		node.links = append(node.links, child)
	} else {
		tree[parent] = &Node{links: []int{child}}
	}
}

func (node *Node) getChildren(parent_id int) []int {
	ret := make([]int, 0)
	for _, v := range node.links {
		if v == parent_id {
			continue
		}
		if _, ok := appear[v]; ok {
			continue
		}

		ret = append(ret, v)
	}
	return ret
}

// 6 7
// 1 2
// 1 3
// 2 3
// 4 2
// 6 5
// 4 6
// 4 5

// 1 l
// 2 l
// 3 l
// 4
// 5
// 6

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
		// pointer
		if dp, ok := v.(map[int]*Node); ok {
			fmt.Print("[ ")
			for i, v := range dp {
				if i == 0 {
					continue
				}
				fmt.Print(v.links)
				fmt.Print(" ")
			}
			fmt.Println("]")
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
				fmt.Print(", ")
			}
			continue
		}
		if dp, ok := v.([][]float64); ok {
			for _, v := range dp {
				fmt.Print(v)
				fmt.Print(", ")
			}
			continue
		}
		fmt.Print(v)
		fmt.Print(", ")
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
