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
	max_int64 = 9223372036854775807
	// prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)
var dp [][]int
var sli []int
var flag [][]int
var sum []int

func run() interface{} {
	n := readInt()
	// s := read()

	sli = readSli(n)
	sum = make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] += sli[i] + sum[i]
	}

	// dp[L][R]
	dp = make([][]int, n)
	flag = make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		flag[i] = make([]int, n)
	}

	ans := calculate(0, n-1)
	return ans
}

func calculate(l, r int) int {
	// p(r+1, l+1)
	if flag[l][r] == 1 {
		return dp[l][r]
	}

	flag[l][r] = 1

	if l == r {
		return 0
	}

	min := max_int64
	for i := l; i < r; i++ {
		// divide point
		div := i
		min = Min(min, calculate(l, div)+calculate(div+1, r))
	}

	// p(Sum(l, r), min)
	dp[l][r] = min + Sum(l, r)

	return dp[l][r]
}

func Sum(l, r int) int {
	// p(sum[r+1], sum[l])
	return sum[r+1] - sum[l]
}

// 3
// 10 10 10
// 10 10

// 1 2 3 4

// c(1,4)
// l,r = 1,4
//   for
//     c(1,1)+c(2,4) = 1 + c(2,4)
//       for c(2,4)
//         [1]= c(2,2)+c(3,4) = 2+7(7) = 9(16)
//         [2]= c(2,3)+c(4,4) = 5(5)+4 = 9(14)
//         c(2,4) = 9(14) dp(1,3)= 14
//     = 1 + 9 = 10
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
