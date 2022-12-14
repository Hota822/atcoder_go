package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
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

type Cn struct {
	c string
	n int
}

func run() {
	// n := readInt()
	short := read()
	long := read()

	var s string
	var t string
	if len(short) < len(long) {
		s = short
		t = long
	} else {
		s = long
		t = short
	}

	// sli := make([][]int, n)
	// for i:=0; i < n; i++ {
	// 	sli[i] = readSli(2)
	// }
	s_sli := strings.Split(s, "")
	t_sli := strings.Split(t, "")

	s_len := len(s_sli)
	t_len := len(t_sli)

	// alp_map := GetAlphabetMap()
	// count := make([]int, 26)
	// for _, v := range s {
	// count[v-97]++
	// }

	dp := make([][]int, t_len+1)
	dp[0] = make([]int, s_len+1)
	for i := 1; i <= t_len; i++ {
		// chr_count := 0
		dp[i] = make([]int, s_len+1)
		for j := 1; j <= s_len; j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = Max(dp[i][j-1], dp[i-1][j-1]+1)
				// chr_count++
			} else {
				// dp[i][j] = dp[i][j-1]
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	ans := make([]string, t_len)

	// max_len := dp[t_len][s_len]
	for i, j := t_len, s_len; i > 0 && j > 0; {
		// if dp[i][j] == max_len {
		// 	// fmt.Print(dp[i][j])
		// 	ans = append(ans, s_sli[j])
		// 	max_len--
		//     i--
		//     j--
		// }
		if dp[i-1][j] == dp[i][j]-1 && dp[i][j-1] == dp[i][j]-1 {
			ans = append(ans, s_sli[j-1])
			i--
			j--
		} else if dp[i-1][j] == dp[i][j]-1 {
			j--
		} else {
			i--
		}
	}

	for i := len(ans); i > 0; i-- {
		fmt.Print(ans[i-1])
	}
	fmt.Println()
}

//   a x y b
// a 1 1 1 1
// b 1 1 1 2
// y 1 1 2 2
// x 1 2 2 2
// b 1 2 2 3

//	a b r a c a d a b r a
//
// a[0 1 1 1 1 1 1 1 1 1 1 1]
// v[0 1 1 1 1 1 1 1 1 1 1 1]
// a[0 1 1 1 2 2 2 2 2 2 2 2]
// d[0 1 1 1 2 2 2 3 3 3 3 3]
// a[0 1 1 1 2 2 3 3 4 4 4 4]
// k[0 1 1 1 2 2 3 3 4 4 4 4]
// e[0 1 1 1 2 2 3 3 4 4 4 4]
// d[0 1 1 1 2 2 3 4 4 4 4 4]
// a[0 1 1 1 2 2 3 4 5 5 5 5]
// v[0 1 1 1 2 2 3 4 5 5 5 5]
// r[0 1 1 2 2 2 3 4 5 5 6 6]
// a[0 1 1 2 3 3 3 4 5 5 6 7]
//
// if match left or (above-left +1)
// if not match, use max of above and left
// start to read from under right of dp
// then output before chr if count down
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

// =======================main========================
func main() {
	buf := make([]byte, initial_buf)
	sc.Buffer(buf, max_bufSize)
	sc.Split(bufio.ScanWords)
	run()
	// print(result)
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
		fmt.Println(v)
	}
}

func pdp(arg ...interface{}) {
	_, _, l, _ := runtime.Caller(1)
	s := strconv.Itoa(l)
	fmt.Println("dumped dp at line: " + s + ", value: ")
	if sli, ok := arg[0].([][]string); ok {
		for _, s := range sli {
			fmt.Println(s)
		}
		return
	}
	if sli, ok := arg[0].([][]int); ok {
		for _, s := range sli {
			fmt.Println(s)
		}
		return
	}

}
