package main

import (
	"fmt"
	// "bufio"
	"os"
	// "strings"
	// "strconv"
	// "reflect"
	// "math"
	// "sort"
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	// var x int
	// fmt.Scan(&x)

	sc.Split(bufio.ScanWords)

	x := readInt()
	// y := readInt()

	// s := read()
	// t := read()

	// n := readInt()
	// sli := make([]int, n)
	// sli := make([]string, n)

	return x
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
