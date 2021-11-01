package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	// "strings"
	"runtime"
	// "math"
	// "reflect"
	"sort"
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
	// s := read()
	els := make([][]int, 3)
	els[0] = []int {1, 2}
	els[1] = []int {0, 2}
	els[2] = []int {0, 1}
	
	dp := make([][]int, n)
	dp[0] = readSli(3)

	for i:=1; i < n; i++ {
		dp[i] = make([]int, 3)
		sli := readSli(3)
		for idx, e := range els {
			max_val := 0
			for _, element := range e {
				max_val = Max(max_val, dp[i -1][element] + sli[idx])
				// d(dp[i -1][element], sli[idx])
			}
			dp[i][idx] = max_val
			// d(max_val)
		}
	}

	ans := MaxSli(dp[n -1])
	return ans
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxSli(sli []int) int {
	sort.Ints(sli)
	l := len(sli)
	return sli[l -1]
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

func d(arg ...interface{}) {
	_, _, l, _ := runtime.Caller(1)
	s := strconv.Itoa(l)
	fmt.Println("dumped at line: " + s + ", value: ")
	for _, v := range arg {
		fmt.Println(v)
	}
}
