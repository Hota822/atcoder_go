package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	// "runtime"
	"math"
	// "reflect"
	// "sort"
	// "strings"
)

const (
    // max_bufSize = 100000 // default: 64000
	// prime_number = 1000000007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	// buf := make([]byte, maxBufSize)
	// sc.Buffer(buf, max_bufSize)
	sc.Split(bufio.ScanWords)

	n := readInt()
	cost := make([]int, n)
	sli := make([]int, n)
	sli[0] = readInt()
	cost[0] = 0
	sli[1] = readInt()
	cost[1] = DiffAbs(sli[0], sli[1])

	for i:=2; i < n; i++ {
		h := readInt()
		sli[i] = h
		c1 := DiffAbs(sli[i -1], h)
		c2 := DiffAbs(sli[i -2], h)
		cost[i] = Min(c1 + cost[i -1], c2 + cost[i -2])
	}
	return cost[n -1]
}


func readInt() int {
    sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
    return ret
}

func DiffAbs(x, y int) int {
	return int(math.Abs(float64(x - y)))
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
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

// func d(arg ...interface{}) {
// 	_, _, l, _ := runtime.Caller(1)
// 	s := strconv.Itoa(l)
// 	fmt.Println("dumped at line: " + s + ", value: ")
// 	for _, v := range arg {
// 		fmt.Println(v)
// 	}
// }
