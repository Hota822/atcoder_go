package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
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
	// n := readInt()
	s := read()

	sli := GetPermutation(s)
	// sli := readSli(n)
	m := make(map[string]struct{})
	for i:=0; i < len(sli); i++ {
		m[sli[i]] = struct{}{}
	}
	ans := len(m)
	return ans
}

func GetPermutation(s string) []string {
	swap := func(i, j int, sli []string) []string {
		new := make([]string, len(sli))
		copy(new, sli)
		new[i], new[j] = sli[j], sli[i]
		return new
	}

	var n int = 1
	l := len(s)
	for i := l; i > 0; i-- {
		n *= i
	}
	if l == 1 {
		return []string {s}
	}

	idx := [2]int {0, 1}
	ret := make([]string, n)
	ret_sli := make([][]string, 0, n)
	ret_sli = append(ret_sli, strings.Split(s, ""))
	for {
		add := make([][]string, 0, (l -1) * len(ret_sli))
		for {
			for _, sli := range ret_sli {
				add = append(add, swap(idx[0], idx[1], sli))
			}
			if idx[1] +1 == l {
				ret_sli = append(ret_sli, add...)
				break
			} else {
				idx[1]++
			}
		}
		if idx[0] +1 == idx[1] {
			break
		} else {
			idx[0]++
			idx[1] = idx[0] + 1
		}
	}
	for pos, sli := range ret_sli {
		ret[pos] = strings.Join(sli, "")
	}
	return ret
}

// ========================read
func read() string {
	sc.Scan()
    return sc.Text()
}

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
