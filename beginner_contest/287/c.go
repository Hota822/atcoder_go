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

var tree map[int]*Node
var root map[int]struct{}

func run() interface{} {
	n, m := readInt(), readInt()
	if m == 0 {
		return false
	}

	tree = make(map[int]*Node)
	root = make(map[int]struct{})
	for i := 0; i < m; i++ {
		v1, v2 := readInt(), readInt()
		createNode(v1, v2)
		createNode(v2, v1)
	}

	if len(root) != 2 {
		return false
	}
	// p(root)

	parent := 0
	current := 0
	for v := range root {
		current = v
		// p(v)
		break
	}
	for i := 1; i < n; i++ {
		// p(i)
		node := tree[current]
		// pl(node)
		children := node.getChildren(parent)

		if len(children) != 1 {
			return false
		}
		parent = current
		current = children[0]
		// pl(current, children)
	}
	last_node := tree[n]
	last := last_node.getChildren(parent)
	// p()
	if len(last) != 0 {
		// p(last_node, current)
		return false
	}

	return true
}

func createNode(x, y int) {
	if node, ok := tree[x]; ok {
		node.links = append(node.links, y)
		delete(root, x)
	} else {
		tree[x] = &Node{links: []int{y}}
		root[x] = struct{}{}
	}
}

func (node *Node) getChildren(parent int) []int {
	ret := make([]int, 0)
	for _, v := range node.links {
		if v == parent {
			continue
		}

		ret = append(ret, v)
	}
	return ret
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
