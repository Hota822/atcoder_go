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

	sli := make([][]int, n+1)
    sum_v := 0
	for i:=1; i <= n; i++ {
		sli[i] = readSli(2) // [weight, value]
	}
    dp := make([][]int, n +1) // dp[value][number] -> min wei

    vm := 100_000
    // vm:=50

    for i:=0; i<=n; i++ {
        dp[i] = make([]int, vm + 1)
    }

    sum_v = 0
    for i:= 1; i<=n; i++ {

        // values
        item := sli[i]
        sum_v += item[1]
        for j:=1; j<=vm; j++ {
            if j < item[1] { dp[i][j] = CustomMin(dp[i -1][j], item[0]); continue }
            if j > sum_v { break}

            dp[i][j] = CustomMin(dp[i -1][j], dp[i -1][j -item[1]] + item[0])
        }
    }

	for i:= vm; i>0; i-- {
        if dp[n][i] <= w && dp[n][i] > 0 {
            return i
        }
    }
    return 0
}

func CustomMin(x, y int) int {
    if x == 0 { return y }
    if y == 0 { return x }

    return Min(x, y)
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
