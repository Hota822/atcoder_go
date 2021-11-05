package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"runtime"
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
	n, w := readInt(), readInt()
	// s := read()
	v_max := 0
	sli := make([][]int, n+1)
	for i:=1; i <= n; i++ {
		sli[i] = readSli(2)
		v_max += sli[i][1]
	}

	dp := make([][]int, n+1)
	dp[0] = make([]int, 100_000 +1)
	v_sum := 0
	for i:=1; i <= n; i++ {
		dp[i] = make([]int, 100_000 +1)
		ns := sli[i]
		v_sum += ns[1]
		for j:=1; j <= 100_000; j++ {
			if ns[1] > j {
				// 重さが少なく、交換する場合
				before := dp[i -1][j]
				dp[i][j] = Min(before, ns[0])
			} else {
				if v_sum < j {
					// 最大値を超えたら終了
					break
				}
				// 入れる or 入れないの最少
				in := dp[i -1][j - ns[1]] + ns[0]
				not_in := dp[i -1][j]
				dp[i][j] = Min(in, not_in)
			}
		}
	}
	
	ans := 0
	for i:=100_000; i > 0; i--{
		if dp[n][i] <= w && dp[n][i] > 0 {
			ans = i
			break
		}
	}
	// ans := sli
	// p(dp[n])
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
	if x == 0 {
		x = 100_000_000_000
	}
	if y == 0 {
		y = 100_000_000_000
	}
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
