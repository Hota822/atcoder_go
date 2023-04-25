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
	prime_number = 1000_000_007
	white        = 0
	black        = 1
)

var sc = bufio.NewScanner(os.Stdin)

// var sli []int
// var memo [][]int
var dp [][2]int
var tree map[int]*Node

type Node struct {
	links              []int
	self               int
	children           []int
	alreadyGetChildren bool
}

// 879053727
// 879053727
// 323768596

func run() interface{} {
	n := readInt()
	// s := read()
	if n == 1 {
		return 2
	}

	tree = make(map[int]*Node)
	dp = make([][2]int, n+1)
	dp[n-1] = [2]int{}
	dp[n] = [2]int{}
	for i := 0; i < n-1; i++ {
		dp[i] = [2]int{}
		x, y := readInt(), readInt()
		createNode(x, y)
		createNode(y, x)
	}

	root := 1
	rootNode := tree[root]
	b := calculate(black, root, rootNode.getChildren(0))
	w := calculate(white, root, rootNode.getChildren(0))
	ans := b + w
	return ans % prime_number
}

func createNode(a, b int) {
	if node, ok := tree[a]; ok {
		node.links = append(node.links, b)
	} else {
		tree[a] = &(Node{links: []int{b}})
	}
}

func (node *Node) getChildren(parentIdx int) []int {
	if node.alreadyGetChildren {
		return node.children
	}

	node.alreadyGetChildren = true
	ret := make([]int, 0)
	for _, idx := range node.links {
		if idx != parentIdx {
			ret = append(ret, idx)
		}
	}
	node.children = ret
	return ret
}

func calculate(parentColor, parentIdx int, children []int) int {
	if len(children) == 0 {
		return 1
	}

	if dp[parentIdx][parentColor] > 0 {
		return dp[parentIdx][parentColor]
	}

	ret := 1
	for _, idx := range children {
		child := tree[idx]
		w := calculate(white, idx, child.getChildren(parentIdx)) % prime_number
		dp[idx][white] = w
		if parentColor == black {
			ret *= w
			ret %= prime_number
			continue
		}

		b := calculate(black, idx, child.getChildren(parentIdx)) % prime_number
		dp[idx][black] = b
		ret *= (w + b)
		ret %= prime_number
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
				fmt.Print(i)
				fmt.Print(":")
				fmt.Print(v.links)
				fmt.Print(", ")
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
