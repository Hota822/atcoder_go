package main

import (
	"bufio"
	"fmt"
	"math/bits"
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
	// max_int64 = 9223372036854775807
	prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	n := readInt()

	sli := make([][]int, n)
	for i := 0; i < n; i++ {
		sli[i] = readSli(n)
	}

	dp := make([]int, 1<<n)
	dp[0] = 1
	// S: マッチした女のビット集合
	for S := 0; S < 1<<n; S++ {
		// S==0の時、マッチしても追加する値が0なのでスキップ可能
		if dp[S] == 0 {
			continue
		}

		// 立っているビット数をカウント > マッチした人数
		// i = 0, 1, 2, 3, 4...
		i := bits.OnesCount(uint(S))
		for j := 0; j < n; j++ {
			// j番目のビットを確認し、既にマッチ済みの時はスキップ
			if S>>j&1 == 1 {
				continue
			}

			// i番目の男とj番目の女が相性がいいとき
			if sli[i][j] == 1 {
				// j番目の女とマッチング（ビットを立てる）
				// そこにマッチ前の場合の数を入れる
				dp[S|1<<j] += dp[S]
				dp[S|1<<j] %= prime_number
			}
		}
	}

	ans := dp[1<<n-1]

	return ans
}

// 3
//i\j
// 0 1 1
// 1 0 1
// 1 1 1
// S==0, into j loop
//   j=0 S>>j == 0 > sli[i][j] == 0, skip
//   j=1 same      > S|1<<j = 10(2) == 2, dp[S|1<<j]++
//   j=2 same      > S|1<<j = 100(2) == 4, dp[S|1<<j]++
//   dp = [1 0 1 0 1 0 0 0]
// S==1, i == 1, dp[S] == 0, continue
// S==2, i == 1, dp[S] == 1, into j loop
//   j=0 S>>j == 10(2)>>0 !=0, > S|1<<j = 11(2) == 3, dp[S|1<<j]++
//   j=1 S>>j == 10(2)>>1 == 1, skip
//   j=2 S>>j == 10(2)>>2 !=0, > S|1<<2 = 110 == 6, dp[S|1<<j]++
//   dp = [1 0 1 1 1 0 1 0]
// S==3, 4, 5 ...

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

func ps(header []string, arg ...interface{}) {
	_, _, l, _ := runtime.Caller(1)
	s := strconv.Itoa(l)
	fmt.Println("dumped at line: " + s + ", value: ")
	for idx, v := range arg {
		fmt.Print(header[idx])
		fmt.Print(": ")
		fmt.Print(v)
		fmt.Print(", ")
	}
	fmt.Println()
}
