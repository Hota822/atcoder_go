package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
	// "strconv"
	// "reflect"
	// "math"
	"sort"
)

const (
    // maxBufSize = 100000
)

// var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	// buf := make([]byte, maxBufSize)
	// sc.Buffer(buf, maxBufSize)
	// sc.Split(bufio.ScanWords)

	// x := readInt()
	// y := readInt()

	x := SortBy {}
	x[3] = 6
	x[1] = 2
	x[2] = 4
	x[100] = 200
	sort.Sort(x)
	for k, v := range x {
		fmt.Println(k, v)
	}
	
	// s := read()
	// t := read()

	// n := readInt()
	// sli := make([]int, n)
	// sli := make([]string, n)

	return x
}

type SortBy map[int]int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }




// ========================read
// func read() string {
// 	sc.Scan()
//     ret := sc.Text()
//     return ret
// }

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
