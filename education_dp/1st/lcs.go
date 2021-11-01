package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"runtime"
	"strings"
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
	// n := readInt()
	s, t := read(), read()

	// sli := make([][]int, n)
	// for i:=0; i < n; i++ {
	// 	sli[i] = readSli(2)
	// }
	t_sli := strings.Split(t, "")
	idx := 0
	ans := make([]string, 0, Max(len(s), len(t)))
	m := make(map[string]struct{}, 3000)
	for _, c := range s {
		for j:=idx; j < len(t); j++ {
			if t_sli[j] == string(c) {
				m[string(c)] = struct{}{}
			}
		}
	}
	// dp := make([][]int, n)

	// ans := sli
	return ans
}


// ========================read
func read() string {
	sc.Scan()
    return sc.Text()
}

// func readSli(n int) []string {
// func readSli(n int) []int {
// 	// sli := make([]string, n)
// 	sli := make([]int, n)
// 	for i:=0; i < n; i++ {
// 		// sli[i] = read()
// 		sli[i] = readInt()
// 	}
//     return sli
// }

// func readInt() int {
//     sc.Scan()
//     ret, _ := strconv.Atoi(sc.Text())
//     return ret
// }


//=======================main========================
func main() {
	buf := make([]byte, initial_buf)
	sc.Buffer(buf, max_bufSize)
	sc.Split(bufio.ScanWords)
	result := run()
	PrintOne(result)
}

func Print(ans interface{}) {
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

func PrintOne(ans interface{}) {
	if sli, ok := ans.([]int); ok {
		for _, v := range sli {
			fmt.Print(v)
		}
	}
	if sli, ok := ans.([]string); ok {
		for _, v := range sli {
			fmt.Print(v)
		}
	}
	fmt.Println("")
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
