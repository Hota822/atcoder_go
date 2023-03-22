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

var dp [][][]int
var sli []int
var memo [][]bool

// var dp [][]int

func run() interface{} {
	n := readInt()
	// s := read()

	if n == 1 {
		return readInt()
	}
	// sli := make([][]int, n)
	// // sli := make([][]int, n)
	// for i := 0; i < n; i++ {
	// 	sli[i] = readSli(2)
	// }
	sli = readSli(n)

	dp = make([][][]int, n+1)
	memo = make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		// dp[i] = make([]int, n+1)
		dp[i] = make([][]int, n+1)
		memo[i] = make([]bool, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]int, 3)
		}
	}

	// l, rのインデックスを添字もつDPを考える
	// 1: first, 2: second
	first, second := calculate(0, n-1, 1)
	ans := first - second
	return ans
}

// return first, second
func calculate(l, r, player int) (int, int) {
	if memo[l][r] {
		return dp[l][r][1], dp[l][r][2]
	}

	memo[l][r] = true
	if r-l == 1 {
		lv, rv := sli[l], sli[r]
		if player == 1 {
			if lv > rv {
				dp[l][r][1] = lv
				dp[l][r][2] = rv
				return lv, rv
			} else {
				dp[l][r][1] = rv
				dp[l][r][2] = lv
				return rv, lv
			}
		} else {
			if lv > rv {
				dp[l][r][1] = rv
				dp[l][r][2] = lv
				return rv, lv
			} else {
				dp[l][r][1] = lv
				dp[l][r][2] = rv
				return lv, rv
			}

		}
	}

	l_first, l_second := calculate(l+1, r, another(player))
	r_first, r_second := calculate(l, r-1, another(player))
	if player == 1 {
		l_first += sli[l]
		r_first += sli[r]
		if l_first > r_first {
			dp[l][r][1] = l_first
			dp[l][r][2] = l_second
			return l_first, l_second
		} else {
			dp[l][r][1] = r_first
			dp[l][r][2] = r_second
			return r_first, r_second
		}
	} else {
		l_second += sli[l]
		r_second += sli[r]
		if l_second > r_second {
			dp[l][r][1] = l_first
			dp[l][r][2] = l_second
			return l_first, l_second
		} else {
			dp[l][r][1] = r_first
			dp[l][r][2] = r_second
			return r_first, r_second
		}
	}
}

// func sort(x, y int) int, int {
//     if x > y {
//         return x
//     } else {
//         return y
//     }
// }

func another(player int) int {
	if player == 1 {
		return 2
	} else {
		return 1
	}
}

// 4
// 10 80 90 30
// cal(0, 3)
//     1: 10 + cal(1,3) or 30 + cal(0,2)
//     2(1,3): 80 + cal(2,3) or 30 + cal(1,2)
//     2(0,2): 10 + cal(1,2) or 90 + cal(0, 1)
//     1:

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
