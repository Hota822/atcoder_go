package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	// "strconv"
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

	x := read()
	l := len(x)
	if l == 1 {
		return "000" + x
	}
	if l == 2 {
		return "00" + x
	}
	if l == 3 {
		return "0" + x
	}
	// y := readInt()

	// s := read()
	// t := read()

	// n := readInt()
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
