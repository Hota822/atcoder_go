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

func run() interface{} {
	n := readInt()
    w := readInt()
	// s := read()

	sli := make([][]int, n)
	for i:=0; i < n; i++ {
		sli[i] = readSli(2) // [weight, value]
	}

    dp:= make([][]int, w + 1)

    dp[0] = make([]int, n + 1,)

    // weight
    for i:=1; i<=w; i++ {
        // into number of n or not
        dp[i] = make([]int, n + 1)
        for j:=1; j<=n; j++ {

            wei:= sli[j -1][0]
            // 入れない
            if i - wei < 0 {
                dp[i][j] = dp[i][j -1]
                continue
            }
            // 入れた時の最大 = j からweiを引いた値に今のvalを足す
            before := dp[i -wei][j -1] + sli[j -1][1]
            dp[i][j] = Max(dp[i][j -1], before)
        }
    }
	ans := dp[w][n]
	return ans
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
	for i:=0; i < n; i++ {
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


//=======================main========================
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
