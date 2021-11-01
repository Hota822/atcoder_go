package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	// "strings"
	// "runtime"
	// "math"
	// "reflect"
	"sort"
)

const (
    max_bufSize = 1_000_000_000 // default: 65536
	initial_buf = 10000
	// max_int32 = 2147483647
	// prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

func run()  {
	t := readInt()
	// s := read()
	for i:=0; i < t; i++ {
		fmt.Println(GetMin(readSli(3)))
	}

	// sli := readSli(n)

	// ans := n
	// return ans
}

func GetMin(sli []int) int {
	sort.Ints(sli)
	base := sli[0]
	sli[2] -= base
	sli[1] -= base
	min := sli[2]

	a_b := sli[1]
	b_c := sli[2] - sli[1]
	a_c := sli[2]

	c, c_ok := GetTimes(a_b)
	b, b_ok := GetTimes(a_c)
	a, _ := GetTimes(b_c)
	if a < 0 && b < 0 && c < 0 {
		return -1
	}

	if c_ok {
		min = Min(min, c)
	}

	if b_ok {
		min = Min(min, b)
	}

	return base + min
}

func GetTimes(abs int) (int, bool) {
	mod := abs / 3
	if mod != 0 {
		return -1, false
	}

	return abs, true
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
	run()
	// print(result)
}

// func print(ans interface{}) {
// 	if v, ok := ans.(bool); ok {
// 		if v {
// 			fmt.Println("Yes")
// 		} else {
// 			fmt.Println("No")
// 		}
// 		return
// 	}
// 	if sli, ok := ans.([]int); ok {
// 		for _, v := range sli {
// 			fmt.Println(v)
// 		}
// 		return
// 	}
// 	if sli, ok := ans.([]int64); ok {
// 		for _, v := range sli {
// 			fmt.Println(v)
// 		}
// 		return
// 	}
// 	if sli, ok := ans.([]string); ok {
// 		for _, v := range sli {
// 			fmt.Println(v)
// 		}
// 		return
// 	}
// 	fmt.Println(ans)
// }

// func d(arg ...interface{}) {
// 	_, _, l, _ := runtime.Caller(1)
// 	s := strconv.Itoa(l)
// 	fmt.Println("dumped at line: " + s + ", value: ")
// 	for _, v := range arg {
// 		fmt.Println(v)
// 	}
// }
