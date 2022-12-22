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
	// prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)
var n int
var dp [][][]float64
var flag [][][]int

func run() interface{} {
	n = readInt()
	// s := read()

	// sli := readSliAsFloat(n)
	sli := make([]int, n)
	c := make([]int, 3)
	for i := 0; i < n; i++ {
		sli[i] = readInt()
		c[sli[i]-1]++
	}

	dp = make([][][]float64, n+1)
	flag = make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][]float64, n+1)
		flag[i] = make([][]int, n+1)
		for j := 0; j <= n; j++ {
			dp[i][j] = make([]float64, n+1)
			flag[i][j] = make([]int, n+1)
		}
	}
	flag[0][0][0] = 1
	dp[0][0][0] = 0.0
	ans := calculate(c[0], c[1], c[2])
	return ans
}

func calculate(c1, c2, c3 int) float64 {
	if flag[c1][c2][c3] == 1 {
		return dp[c1][c2][c3]
	}

	flag[c1][c2][c3] = 1

	var e1 float64
	var e2 float64
	var e3 float64

	// if c == n {
	// 	return 1.0
	// }

	p0 := float64(n-c1-c2-c3) / float64(n)
	if c1 > 0 {
		p1 := float64(c1) / float64(n)
		e1 = calculate(c1-1, c2, c3) * p1
	} else {
		e1 = 0.0
	}

	if c2 > 0 {
		p2 := float64(c2) / float64(n)
		e2 = calculate(c1+1, c2-1, c3) * p2
	} else {
		e2 = 0.0
	}

	if c3 > 0 {
		p3 := float64(c3) / float64(n)
		e3 = calculate(c1, c2+1, c3-1) * p3
	} else {
		e3 = 0.0
	}

	e := (1 + e1 + e2 + e3) / (1.0 - p0)
	// p("result", []int{c1, c2, c3}, []float64{e1, e2, e3, e, 1.0 - p0})
	dp[c1][c2][c3] = e
	return e
}

// p0, p1, p2, p3 = (n-c1-c2-c3)/n, c1/n, c2/n, c3/n,
// dp[c1][c2][c3] = （減った後期待値）＋（今回分の期待値）
//  (減った後の期待値) =
//   dp[c1-1][c2][c3]*p1
//   + dp[c1+1][c2-1][c3]*p2
//   + d[c1][c2+1][c3-1]*p3
//  = e1 + e2 + e3

//  (今回分の期待値) = (ゼロが選ばれる期待値) + （ゼロでない値が選ばれる期待値）
//  = dp[c1][c2][c3]*p0 + 1
//

// dp[c1][c2][c3](1-p0) = e1 + e2 + e3 +1
// dp[c1][c2][c3] = (e1 + e2 + e3 + 1) / (1-p0)

// ========================read
// func read() string {
// 	sc.Scan()
//     return sc.Text()
// }

// func readSli(n int) []string {

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
