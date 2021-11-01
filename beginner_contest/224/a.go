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
	s := read()
	l := len(s)
	er := s[(l -2):]
	if er == "er" {
		return "er"
	}
	

	// if ist == "ist" {
		return "ist"
	// }
	// ret
	// sli := readSli(n)

	// ans := n
	// return ans
}


// ========================read
func read() string {
	sc.Scan()
    return sc.Text()
}

// func readSli(n int) []string {
// // func readSli(n int) []int {
// 	sli := make([]string, n)
// 	// sli := make([]int, n)
// 	for i:=0; i < n; i++ {
// 		sli[i] = read()
// 		// sli[i] = readInt()
// 	}
//     return sli
// }

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

// func d(arg ...interface{}) {
// 	_, _, l, _ := runtime.Caller(1)
// 	s := strconv.Itoa(l)
// 	fmt.Println("dumped at line: " + s + ", value: ")
// 	for _, v := range arg {
// 		fmt.Println(v)
// 	}
// }
