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
	n, k := readInt(), readInt()
	// s := read()

	sli := readSli(n)
	
	dp := make([]int, n)
	
	dp[0] = 0
	for i:=1; i < n; i++ {
		min_cost := 1_000_000_000_000
		for j:=1; j <= k; j++ {
			if i -j < 0 {
				break
			}
			step := Abs(sli[i] - sli[i -j]) + dp[i -j]
			// d(step)
			min_cost = Min(min_cost, step)
		}
		// d("end")
		// d(min_cost)
		dp[i] += min_cost
	}

	ans := dp[n -1]
	return ans
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return - x
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
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
