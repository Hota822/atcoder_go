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
var memo map[int]int

func run() interface{} {
	n := readInt()
	// s := read()

	memo = make(map[int]int)

	ans := 0
	factorization(n - 1)
	for i := 1; i < n; i++ {
		ab := i
		cd := n - i
		ans += memo[ab] * memo[cd]
	}

	return ans
}

func factorization(n int) {
	prime_numbers := []int{2, 3}
	memo[1] = 1
	memo[2] = 2
	memo[3] = 2
	for i := 4; i <= n; i++ {
		if memo[i] > 0 {
			continue
		}

		x := i
		m := make(map[int]int)
		is_prime := true
		for j := 0; j < len(prime_numbers); j++ {
			pri := prime_numbers[j]
			for {
				if x%pri == 0 {
					x /= pri
					m[pri]++
					is_prime = false
				} else {
					break
				}
			}
		}

		if is_prime {
			prime_numbers = append(prime_numbers, i)
			memo[i] = 2
			for j := 0; j < len(prime_numbers)-1; j++ {
				pri := prime_numbers[j]
				if i*pri > n {
					break
				}

				memo[i*pri] = memo[i] * 2
			}
		} else {
			ret := 1
			for _, c := range m {
				ret *= (c + 1)
			}
			memo[i] = ret
		}

	}
}

// 5
// 1,5 5,1
// 10 = 5 * 2 1+1 * 1+1 = 4
// 1,10 2,5 5,2 10,1
// 20 = 5 *2^2 = 1+1 * 2+1 = 6
// 1,20 2,10 4,5 5,4 10,2 20,1
// 40 = 5 *2^3 = 1+1 * 3+1 = 8
// 1,40 2,20 4,10 5,8

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