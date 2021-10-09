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
	// sli := make([]int, n)
	// sli := make([]string, n)
	m := make(map[int]int, n)
	var ans int64 = 1
	for i:=0; i < n; i++ {
		c := readInt()
		if _, ok := m[c]; ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}
	var used int = 0
	for k, v := range m {
		fac := permutation(k - used, v)
		used += v
		ans = (ans * fac) % 1000000007
	}
	return ans
}

func permutation(n, v int) int64 {
	var ret int64 = 1
	for i:=0; i < v; i++ {
		ret = (ret * int64(n)) % 1000000007
		n--
		fmt.Println(ret)
	}
	if ret > 0 {
		return ret
	}
	return 0
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