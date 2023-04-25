package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	// "strings"
	// "math"
	// "reflect"
	// "sort"
)

const (
	max_bufSize = 1_000_000_000 // default: 65536
	initial_buf = 10000
	// max_int32 = 2147483647
	// max_int64 = 9223372036854775807
	// prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

// var dp [][]int
// var sli []int
var memo []int
var prime_numbers []int

var y int

func run() interface{} {
	n := readInt()
	// s := read()

	memo = make([]int, n)
	prime_numbers = make([]int, 2)
	prime_numbers[0] = 2
	prime_numbers[1] = 3

	ans := 0
	// for i := 1; i < n; i++ {
	// 	ab := i
	// 	cd := n - i
	// 	ans += calculate(ab) * calculate(cd)
	// }
	y = 100000
	for i := 3; i <= y; i++ {
		// if i == 9 {
		// 	p(prime_numbers)
		// }
		ans = soinsuubunkai(i)
	}
	// ans = soinsuubunkai(100)

	// ans := sli
	p(prime_numbers)
	return ans
}

func calculate(x int) int {
	if memo[x] > 0 {
		return memo[x]
	}
	return 1
}

func soinsuubunkai(original_n int) int {
	// sqrt_x := int(math.Sqrt(float64(original_n)))
	m := make(map[int]int)
	n := original_n

	for i := 0; i < len(prime_numbers); i++ {
		count := 0
		pri := prime_numbers[i]
		// if pri > sqrt_x {
		// 	break
		// }

		for {
			if n%pri == 0 {
				n = n / pri
				count++
			} else {
				m[pri] = count
				break
			}
		}

	}
	if n == 1 {
		// warikireru
		ret := 1
		for _, v := range m {
			ret *= v
		}
		p(m)
		return ret
	} else {
		// prime number
		if prime_numbers[len(prime_numbers)-1] < original_n {
			prime_numbers = append(prime_numbers, original_n)
		}
		return 2
	}
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
	for i := 0; i < n; i++ {
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

func readFloat() float64 {
	sc.Scan()
	ret, _ := strconv.ParseFloat(sc.Text(), 64)
	return ret
}

// =======================main========================
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
	return -x
}

func p(arg ...interface{}) {
	_, _, l, _ := runtime.Caller(1)
	s := strconv.Itoa(l)
	fmt.Println("dumped at line: " + s + ", value: ")
	for _, v := range arg {
		if dp, ok := v.([][]int); ok {
			for _, v := range dp {
				fmt.Println(v)
			}
			continue
		}
		if dp, ok := v.([][]float64); ok {
			for _, v := range dp {
				fmt.Println(v)
			}
			continue
		}
		// pointer
		// if dp, ok := v.([]*Rope); ok {
		// 	fmt.Print("[ ")
		// 	for i, v := range dp {
		// 		if i == 0 {
		// 			continue
		// 		}
		// 		fmt.Print(*v)
		// 		fmt.Print(" ")
		// 	}
		// 	fmt.Println("]")
		// 	continue
		// }
		fmt.Println(v)
	}
}

func pr(arg ...interface{}) {
	for _, v := range arg {
		if dp, ok := v.([][]int); ok {
			for _, v := range dp {
				fmt.Print(v)
				fmt.Print(", ")
			}
			continue
		}
		if dp, ok := v.([][]float64); ok {
			for _, v := range dp {
				fmt.Print(v)
				fmt.Print(", ")
			}
			continue
		}
		fmt.Print(v)
		fmt.Print(", ")
	}
	fmt.Println()
}

func pl(arg ...interface{}) {
	fmt.Println()

	pr(arg...)
}

func ps(header []string, arg ...interface{}) {
	_, _, l, _ := runtime.Caller(1)
	s := strconv.Itoa(l)
	fmt.Println("dumped at line: " + s + ", value: ")
	for idx, v := range arg {
		fmt.Print(header[idx])
		fmt.Print(": ")
		fmt.Print(v)
		fmt.Print(", ")
	}
	fmt.Println()
}

func s(args ...interface{}) []string {
	var s_sli []string
	for _, v := range args {
		if s, ok := v.(string); ok {
			s_sli = append(s_sli, s)
		}
	}
	return s_sli
}
