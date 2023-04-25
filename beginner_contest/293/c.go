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
// var memo [][]int

func run() interface{} {
	h, w := readInt(), readInt()
	// s := read()

	sli := make([][]int, h)
	for i := 0; i < h; i++ {
		sli[i] = readSli(w)
	}

	paths := makePaths(h, w)
	// p(paths)
	ans := 0
	for _, v := range paths {
		through := make(map[int]struct{})
		dup := false
		for _, coordinate := range v {
			a := sli[coordinate[0]-1][coordinate[1]-1]
			if _, ok := through[a]; ok {
				dup = true
				break
			} else {
				through[a] = struct{}{}
			}
		}

		if !dup {
			ans++
		}
	}

	return ans

}

func makePaths(x, y int) [][][2]int {
	paths := make([][][2]int, 1)
	paths[0] = make([][2]int, 1)
	paths[0][0] = [2]int{1, 1}

	for i := 1; i < x+y-1; i++ {
		next_paths := make([][][2]int, 0)
		for _, v := range paths {
			last_coordinate := v[i-1]
			if last_coordinate[0] < x {
				var path_to_right [][2]int
				path_to_right = append(path_to_right, v...)
				path_to_right = append(path_to_right, [2]int{last_coordinate[0] + 1, last_coordinate[1]})
				next_paths = append(next_paths, path_to_right)
			}

			if last_coordinate[1] < y {
				var path_to_bottom [][2]int
				path_to_bottom = append(path_to_bottom, v...)
				path_to_bottom = append(path_to_bottom, [2]int{last_coordinate[0], last_coordinate[1] + 1})
				next_paths = append(next_paths, path_to_bottom)
			}
		}
		paths = next_paths
		// p(next_paths)
	}
	return paths
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
		if dp, ok := v.([][][2]int); ok {
			for _, v := range dp {
				fmt.Println(v)
			}
			continue
		}
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
		if dp, ok := v.([][2]int); ok {
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
