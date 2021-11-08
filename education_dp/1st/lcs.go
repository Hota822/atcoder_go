package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"runtime"
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

func run() interface{} {
	// n := readInt()
	s, t := read(), read()

	// sli := make([][]int, n)
	// for i:=0; i < n; i++ {
	// 	sli[i] = readSli(2)
	// }
	s_sli := strings.Split(s, "")
	t_sli := strings.Split(t, "")
	dp := make([][]int, len(s) +1)
    dp[0] = make([]int, len(t) +1)

    for i:=1; i <= len(s); i++ {
        s_c := s_sli[i -1]
        dp[i] = make([]int, len(t) +1)
        for j:=1; j <=len(t); j++ {
            t_c := t_sli[j -1]
            if s_c == t_c {
                dp[i][j] = dp[i -1][j -1] +1
            } else {
                dp[i][j] = Max(dp[i -1][j], dp[i][j -1])
            }
        }
    }

    l := dp[len(s)][len(t)]
    ans := make([]string, l)
    i := len(s)
    j := len(t)
    for {
        s_c := s_sli[i -1]
        t_c := t_sli[j -1]
        if t_c == s_c {
            ans[l -1] = t_c
            i--
            j--
            l--
        } else if dp[i -1][j] == dp[i][j] {
            i--
        } else {
            j--
        }
        if l == 0 {
            break
        }
    }
    // p(dp)
    return ans
}

// ========================read
func read() string {
	sc.Scan()
    return sc.Text()
}

// func readSli(n int) []string {
// func readSli(n int) []int {
// 	// sli := make([]string, n)
// 	sli := make([]int, n)
// 	for i:=0; i < n; i++ {
// 		// sli[i] = read()
// 		sli[i] = readInt()
// 	}
//     return sli
// }

// func readInt() int {
//     sc.Scan()
//     ret, _ := strconv.Atoi(sc.Text())
//     return ret
// }


//=======================main========================
func main() {
	buf := make([]byte, initial_buf)
	sc.Buffer(buf, max_bufSize)
	sc.Split(bufio.ScanWords)
	result := run()
	PrintOne(result)
}

func Print(ans interface{}) {
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

func PrintOne(ans interface{}) {
	if sli, ok := ans.([]int); ok {
		for _, v := range sli {
			fmt.Print(v)
		}
	}
	if sli, ok := ans.([]string); ok {
		for _, v := range sli {
			fmt.Print(v)
		}
	}
	fmt.Println("")
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
	return - x
}

func p(arg ...interface{}) {
	_, _, l, _ := runtime.Caller(1)
	s := strconv.Itoa(l)
	fmt.Println("dumped at line: " + s + ", value: ")
	for _, v := range arg {
		fmt.Println(v)
	}
}
