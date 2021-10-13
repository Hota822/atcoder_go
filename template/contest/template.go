package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	// "strconv"
	// "runtime"
	// "math"
	// "reflect"
	// "sort"
)

const (
    // max_bufSize = 100000 // default: 65536
	// max_int = 2147483647
	// prime_number = 1000000007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	// buf := make([]byte, 0)
	// sc.Buffer(buf, max_bufSize)
	// sc.Split(bufio.ScanWords)

	// x := readInt()
	// s := read()

	// sli := make([]int, n)
	// sli := make([]string, n)

	ans := 1
	return ans
}


// ========================read
func read() (string...) {
	sc.Scan()
    s := sc.Text()
	return strings.Split(s)
}

func read() string {
	sc.Scan()
    return sc.Text()
}

// func readInt() int {
//     sc.Scan()
//     ret, _ := strconv.Atoi(sc.Text())
//     return ret
// }


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
