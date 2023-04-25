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

type Rope struct {
	l_edge int
	r_edge int
}

func run() interface{} {
	n, m := readInt(), readInt()
	// s := read()

	ropes := make([]*Rope, n+1)
	for i := 1; i <= n; i++ {
		r := Rope{l_edge: i, r_edge: i}
		ropes[i] = &r
	}

	line, ring := n, 0
	for i := 0; i < m; i++ {
		a, _, c, _ := readInt(), read(), readInt(), read()
		line--
		rope_a := ropes[a]
		rope_c := ropes[c]

		if rope_a == rope_c {
			// p(ropes)
			ring++
			continue
		}

		rope_a.add(a, rope_c, c)
		// rope_a.l_edge = 1
		l_edge := rope_c.l_edge
		r_edge := rope_c.r_edge
		ropes[l_edge] = ropes[a]
		ropes[r_edge] = ropes[a]
	}

	return []int{ring, line}
}

func (base *Rope) add(from int, add *Rope, to int) {
	var edge int

	if add.l_edge == to {
		edge = add.r_edge
	} else {
		edge = add.l_edge
	}

	if base.l_edge == from {
		base.l_edge = edge
	} else {
		base.r_edge = edge
	}
}

// 5 3
// 3 R 5 B
// 5 R 3 B
// 4 R 2 B

//  1 R B
//  2 R B
//  3 5 B
//  4 R B
//  5 R 3

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
		if dp, ok := v.([]*Rope); ok {
			fmt.Print("[ ")
			for i, v := range dp {
				if i == 0 {
					continue
				}
				fmt.Print(*v)
				fmt.Print(" ")
			}
			fmt.Println("]")
			fmt.Println(v)
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
