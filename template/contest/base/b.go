package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	// "runtime"
	// "math"
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

	x := readInt()
	s := read()

	// sli := make([]int, n)
	// sli := make([]string, n)

	return x
}


// ========================read
func read() string {
	sc.Scan()
    ret := sc.Text()
    return ret
}

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

// func d(arg ...interface{}) {
// 	_, _, l, _ := runtime.Caller(1)
// 	s := strconv.Itoa(l)
// 	fmt.Println("dumped at line: " + s + ", value: ")
// 	for _, v := range arg {
// 		fmt.Println(v)
// 	}
// }
