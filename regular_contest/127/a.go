package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	"strconv"
	// "reflect"
	"math"
	// "sort"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt64()
	len := GetLen(n)
	if len == 1 {
		fmt.Println("ssssss")
		return
	}

	var ans int64 = 0
	// sli := make([]int64, n)
	var sum [15]int64
	sum[0] = 1
	for i := 1; i < 15; i++ {
		sum[i] = sum[i -1] + int64(math.Pow(10, float64(i)))
	}

	// n - 1　桁までの合計を取得
	for i := 0; i < len -1; i++ {
		if len == i {
			break
		}
		ans += sum[i]
	}

	// n 桁目の合計を取得
	if n < sum[int(len) -1] {
		left := makeDelta(len, 1)
		ans += GetLeft(n, left)
	} else if n == sum[int(len) -1] {
		fmt.Println(len)
		left := makeDelta(len, 1)
		ans += GetLeft(n, left)
		ans += int64(len)
	} else {
		ans += GetRight(n, sum[int(len) - 1])
	}
	fmt.Println(ans)
	// fmt.Println(t)
}

func GetLen(n int64) int {
	var i int64 = 1
	var count int64
	for count = 1; count < 16; count++ {
		i *= 10
		result := n / i
		// fmt.Println(result < 1)
		if result < 1 {
			break
		}
	}
	return int(count)
}

func GetLeft(n int64, left []int64) int64 {
	g := GetGravity(n)
	var sum int64 = 0
	for gra, count := range left {
		if gra == g {
			break
		}
		sum += (int64(g +1) * count)
	}
	num := strconv.FormatInt(n, 10)
	new_num := num[int(g + 1):]
	num_as_i, _ := strconv.ParseInt(new_num, 10, 64)
	return sum + num_as_i + 1
}

func GetRight(n, top int64) int64 {
	var sum int64 = 0
	for i := top; i <= n; i++ {
		num := strconv.FormatInt(i, 10)
		var idx int64 = 0
		for _, c := range num {
			if c != '1' {
				break
			}
			idx++
		}
		sum += idx
	}
	return sum
}

func GetGravity(n int64) int {
	num := strconv.FormatInt(n, 10)
	g := 0
	// x := 0
	for _, c := range num {
		if c != '1' {
			// x, _ = strconv.Atoi(string(c))
			break
		}
		g++
	}
	return g
}

func makeDelta(len, base int) []int64 {
	delta := make([]int64, len)
	var i int64 = 1
	var idx int
	for idx = len - 1; idx == 0; idx-- {
		delta[idx] = i * int64(base)
	}
	return delta
}


// func next() string {
// 	sc.Scan()
//     ret := sc.Text()
//     return ret
// }

// func nextInt() int {
//     sc.Scan()
//     ret, _ := strconv.Atoi(sc.Text())
//     return ret
// }

func nextInt64() int64 {
	var ret int64
    sc.Scan()
    ret, _ = strconv.ParseInt(sc.Text(), 10, 64)
    return ret
}

