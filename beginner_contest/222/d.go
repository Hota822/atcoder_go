package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	"strconv"
	// "reflect"
	// "math"
	// "sort"
	"runtime"
)

const (
    // maxBufSize = 100000
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	// buf := make([]byte, maxBufSize)
	// sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	
	// n := readInt()
	a_sli := make([]int, n)
	b_sli := make([]int, n)
	// dp := make([][]int, n)
	for i:=0; i < n; i++ {
		a_sli[i] = readInt()
	}
	for i:=0; i < n; i++ {
		b_sli[i] = readInt()
	}

	diff := b_sli[n -1] - a_sli[n -1] + 1
	// min := [...]int {a_sli[n -1], 0}
	min := a_sli[n -1]
	// fmt.Println(diff)
	ans_sli := make([]int, 0, n)

	for i := a_sli[n -1]; i <= b_sli[n -1]; i++ {
		ans_sli = append(ans_sli, 1)
	}
	// D(a_sli[n -1], b_sli[n -1])
	

	for i:=n -2; i >= 0; i-- {
		before := make([]int, len(ans_sli))
		copy(before, ans_sli)
		ans_sli = make([]int, 0, n)
		idx := 1
		// D(i)
		// D(a_sli[i])
		// D(b_sli[i])
		// D(min, diff)
		// D(ans_sli)
		// D(before)
		// min = a_sli[i]
		for j:=a_sli[i]; j <= b_sli[i]; j++ {
			// D(ans_sli)
			min = Max(min, idx)
			D(min)
			if min <= j {
			// ipf min[0] <= j {
				// ans_sli = append(ans_sli, diff)
			} else {
				D(diff, before)
				diff -= before[idx]
				// ans_sli = append(ans_sli, diff)
				idx++
			}
			ans_sli = append(ans_sli, diff)
			diff = Sum(ans_sli)
		}
		D(ans_sli)
	}
	return diff
}

func Sum(sli []int) int {
	sum := 0
	for _, i:= range sli {
		sum += i
	}
	return sum
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// ========================read
// func read() string {
// 	sc.Scan()
//     ret := sc.Text()
//     return ret
// }

func readInt() int {
    sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
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

func D(arg ...interface{}) {
	_, _, l, _ := runtime.Caller(1)
	s := strconv.Itoa(l)
	fmt.Println("dumped at line: " + s + ", value: ")
	for _, v := range arg {
		fmt.Println(v)
	}
}
