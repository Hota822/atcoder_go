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
	n, w := readInt(), readInt()
	// s := read()

	sli := make([][]int, n +1)
	for i:=1; i <= n; i++ {
		sli[i] = readSli(2)
	}
	
	dp := make([][]int, w +1)
	dp[0] = make([]int, n +1)
	
	// 重さ
	for i:=1; i <= w; i++ {
		// N番目
		dp[i] = make([]int, n +1)
		for j:=1; j <=n; j++ {
			ns := sli[j]
			// d(ns)
			// N番目の重さがi以下の時
			if i < ns[0] {
				// 重すぎる時はそもそも入れられない
					dp[i][j] = dp[i][j -1]
				continue
			} else {
				if i - ns[0] < 0 {
					// 入れると必ずオーバーする場合、１つ前をみて交換するか決める
					before := dp[i][j -1]
					dp[i][j] = Max(before, ns[1])
				} else {
					// １つ前をみて、入れるか入れないかを決める
					// max_val := 0
					not_in := dp[i][j -1]
					in := dp[i -ns[0]][j -1] + ns[1]
					dp[i][j] = Max(not_in, in)
				}
			}
		}
	}
	// d(dp)
	ans := dp[w][n]
	return ans
}

// ============================math
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
