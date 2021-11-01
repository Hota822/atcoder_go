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
	
	// w_max := 100_000_000_000
	v_max := 0
	sli := make([][]int, n +1)
	for i:=1; i <= n; i++ {
		sli[i] = readSli(2)
		v_max += sli[i][1]
	}
	
	dp := make([][]int, n +1)
	for i:=0; i <= n; i++ {
		dp[i] = make([]int, 100001)
		// fmt.Println(len(dp[i]))
		// fmt.Println(dp[i][1])
	}
	
	// fmt.Println(n)
	// number n
	v_sum := 0
	for i:=1; i <= n; i++ {
		// value
		
		kn := sli[i]
		v_sum += kn[1]
		for j:=1; j <= 100000; j++ {
			if v_sum < j {
				break
			}
			if kn[1] > j {
				// j 以下の価値の時は、詰めない時と比較して、重さの低いほうになる
				dp[i][j] = Min(dp[i -1][j], kn[0])
			} else {
				// j以上の時、現在のやつを足したとき、他さないときの最少をとる
				// fmt.Println(kn)
				// fmt.Println(j)
				// fmt.Println(dp[i -1][j - kn[1]])
				dp[i][j] = Min(dp[i -1][j], dp[i -1][j -kn[1]] + kn[0])
				// d(dp[i][j])
			}
		}
	}

	// fmt.Println(dp[n])
	for i:=v_max; i > 0; i-- {
		
		if dp[n][i] <= w  && dp[n][i] > 0 {
			return i
		}
	}

	// d((dp[1]))
	// fmt.Println(dp)
	ans := false
	return ans
}

func Min(x, y int) int {
	if x == 0 {
		x = 100_000_000_000
	}
	if y == 0 {
		y = 100_000_000_000
	}
	if x > y {
		return y
	}
	return x
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
	print(result, false)
}

func print(ans interface{}, debug bool) {
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
	if debug {
		d("end function")
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
