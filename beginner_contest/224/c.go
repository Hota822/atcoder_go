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
	// 1234567890123456789
	// 9223372036854775808
	// max_int32 = 2147483647
	// prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	n := readInt()
	// s := read()
	sli := make([][]int, n)
	for i:=0; i < n; i++ {
		sli[i] = readSli(2)
	}
	// sli := readSli(n)

	ans := 0
	for i:=0; i < n; i++ {
		for j:=i +1; j < n; j++ {
			for k := j +1; k < n; k ++{
				y_zero := false
				x_zero := false
				
				ax, ay := sli[i][0], sli[i][1]
				bx, by := sli[j][0], sli[j][1]
				diff_ab_x := float64(bx) - float64(ax)
				diff_ab_y := float64(by) - float64(ay)
				var gra_ab float64 = 0

				// fmt.Println(diff_ab_y)
				if diff_ab_y == 0 {
					y_zero = true
				} else if diff_ab_x == 0 {
					x_zero = true 
				} else {
					gra_ab = diff_ab_x / diff_ab_y
				}

				cx, cy := sli[k][0], sli[k][1]
				diff_ac_x := float64(cx) - float64(ax)
				diff_ac_y := float64(cy) - float64(ay)
				var gra_ac float64 = 0
				if diff_ac_y == 0 {
					if ! y_zero {
						ans ++
						// fmt.Println("zero")
						continue
					}
				} else if diff_ac_x == 0 {
					if ! x_zero {
						ans ++
						// fmt.Println("zero")
						continue
					}
				} else {
					gra_ac = diff_ac_x / diff_ac_y
					// fmt.Println(gra_ac)
				}

				if gra_ab != gra_ac {
					// if ! y_zero {
						ans++
					// }
					// fmt.Println(ax, ay)
					// fmt.Println(bx, by)
					// fmt.Println(cx, cy)
					// fmt.Println("++", gra_ab, gra_ac)
				}
				// if i == 0 && j == 1 && k == 3 {
				// 	fmt.Println(i, j, k)
				// 	fmt.Println(ax, ay)
				// 	fmt.Println(bx, by)
				// 	fmt.Println(cx, cy)
				// 	fmt.Println(diff_ab_x, diff_ab_y)
				// 	fmt.Println(diff_ac_x, diff_ac_y)
				// }

			}
		}
	}

	// ans := n
	return ans
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
