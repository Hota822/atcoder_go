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
// var masu map[int]bool
// var count int
var memo map[int]*Memo
var x int

type Memo struct {
	masu map[int]bool
	ks   []int
	x    int
}

// var ans int
func getMemo(n int) *Memo {
	if m, ok := memo[n]; ok {
		return m
	} else {
		m := Memo{masu: make(map[int]bool), ks: []int{}, x: 0}
		m.masu[0] = true
		memo[n] = &m
		return &m
	}
}

func run() interface{} {
	t := readInt()
	// s := read()
	memo = make(map[int]*Memo)

	ans_sli := make([]int, 0)
	for i := 0; i < t; i++ {
		n, d, k := readInt(), readInt(), readInt()
		if k == 1 {
			ans_sli = append(ans_sli, 0)
			continue
		}

		m := getMemo(n)
		if len(m.ks) >= k {
			ans_sli = append(ans_sli, m.ks[k-1])
			continue
		}

		x = m.x
		for j := 0; j < n-1; j++ {
			continueNext := calculate(n, d, k, m)
			if !continueNext {
				break
			}
		}
		ans_sli = append(ans_sli, x)
	}

	return ans_sli
}

func calculate(n, d, k int, m *Memo) bool {
	a := x
	x = (a + d) % n
	for {
		if _, ok := m.masu[x]; ok {
			x = (x + 1) % n
		} else {
			if len(m.ks) >= k {
				// ans = x
				return false
			}
			m.masu[x] = true
			m.x = x
			m.ks = append(m.ks, x)

			return true
		}
	}
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
		// if dp, ok := v.([]*Rope); ok {
		// 	fmt.Print("[ ")
		// 	for i, v := range dp {
		// 		if i == 0 {
		// 			continue
		// 		}
		// 		fmt.Print(*v)
		// 		fmt.Print(" ")
		// 	}
		// 	fmt.Println("]")
		// 	continue
		// }
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
