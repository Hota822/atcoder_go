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
	n, q := readInt(), readInt()

	c := make(map[int]Chase)
	sli := make([][]int, q)
	for i:=0; i < q; i++ {
		num := readInt()
		if num == 3 {
			x := readInt()
			sli[i] = []int {3, x}
			c[x] = Chase {0, 0}
		}
		} else {
			x, y := readInt(), readInt()
			sli[i] = []int {num, x, y}
		}
	}

	m := make(map[int][]int)
	for i:=0; i < q; i++ {
		if num == 1 {
			xy := sli[i]
			x, y := xy[0], xy[1]
			if len(m[x]) == 0 {
				m[x] = []int {x, y}
			} else {
				m[x] = append(m[x], y)
			}
			for k, ch := range c {
				if k != 
			}
			continue
		}
		if num == 2 {
			xy := sli[i]
			x, y := xy[0], xy[1]


			continue
		}
		x := readInt()
		Print(m[x])
	}
	// s := read()

	// sli := readSli(n)

	ans := n
	return ans
}

type Chase {
	index int
	head int
}

// ========================read
// func read() string {
// 	sc.Scan()
//     return sc.Text()
// }

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
	run()
	// print(result)
}


func Print(ans interface{}) {
	if v, ok := ans.([]int); ok {
		for _, v := range ans {
			fmt.Print(v)
		}
	}
	if v, ok := ans.([]string); ok {
		for _, v := range ans {
			fmt.Print(v)
		}
	}
	fmt.Println("")
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
