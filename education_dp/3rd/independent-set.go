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
)

var sc = bufio.NewScanner(os.Stdin)
var tree map[int]*Node

type Node struct {
	links []int
}

var dp [][]int
var memo []int

func run() interface{} {
	n := readInt()
	// s := read()

	if n == 1 {
		return 2
	}
	memo = make([]int, n+1)
	makeTree(n)

	dp = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}

	calculate(1, 0)

	ans := (dp[1][0] + dp[1][1]) % prime_number
	return ans
}

func calculate(root int, parent int) {
	if memo[root] == 1 {
		p()
		return
	}
	memo[root] = 1

	dp[root][0] = 1
	dp[root][1] = 1
	// ps([]string{"root", "links"}, root, tree[root].links)
	// p(dp)
	for _, leaf := range tree[root].links {
		if leaf == root || leaf == parent {
			continue
		}

		calculate(leaf, root)
		// for white
		dp[root][0] = dp[root][0] * (dp[leaf][0] + dp[leaf][1]) % prime_number
		// for black
		dp[root][1] = dp[root][1] * dp[leaf][0] % prime_number
	}
}

func makeTree(n int) {
	tree = map[int]*Node{}
	// tree_count := make([]int, n+1)

	createNode := func(x, y int) {
		if _, ok := tree[x]; ok {
			tree[x].links = append(tree[x].links, y)
		} else {
			tree[x] = &Node{links: []int{y}}
		}
	}
	for i := 0; i < n-1; i++ {
		x, y := readInt(), readInt()
		createNode(x, y)
		createNode(y, x)
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
