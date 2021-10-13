package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

const (
    max_bufSize = 10_000_005 // default: 64_000
	// prime_number = 1_000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	buf := make([]byte, max_bufSize)
	sc.Buffer(buf, max_bufSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	dp := make([][3]int, n + 1)
	for i := 0; i < 3; i++ {
		dp[0] = [3]int {0, 0, 0}
	}
	for i:=1; i <= n; i++ {
		abc := [3]int {readInt(), readInt(), readInt()}
		dp[i][0] = MaxOfOthers(dp[i -1], 0) + abc[0]

		dp[i][1] = MaxOfOthers(dp[i -1], 1) + abc[1]

		dp[i][2] = MaxOfOthers(dp[i -1], 2) + abc[2]
	}

	ans := MaxOf3([3]int {dp[n][0], dp[n][1], dp[n][2]})
	return ans
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxOf3(sli [3]int) int {
	x, y, z := sli[0], sli[1], sli[2]
	if x > y {
		if x > z {
			return x
		}
	} else {
		if y > z {
			return y
		}
	}
	return z
}

func MaxOfOthers(sli [3]int, i int) int {
	if i == 0 {
		return Max(sli[1], sli[2])
	} else if i == 1 {
		return Max(sli[0], sli[2])
	}
	return Max(sli[0], sli[1])
}

// ========================read
func readInt() int {
    b := sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
	if !b {
        panic(b)
    }
    return ret
}


//=======================main========================
func main() {
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
