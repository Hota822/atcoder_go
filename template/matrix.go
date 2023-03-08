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

var n int

/*
at first, set length of matrix to n

below is available methods
exponent()     m: multiple times, then read base matrix
multiply()     m: operand, then read matrix. arg is select operator, true: div false: mul
plusWithRead() read 2 matrixes. arg is select operator, true: minus, false: plus
plus()         minus: select operator same as forward, sli1, sli2 : operands
getReverse()   read matrix
getIdentityMatrix()  read n: length of matrix
*/
func run() interface{} {
	n = readInt()

	// iMat := getIdentityMatrix(n)
	// sli := readMatrix()
	// z := multiply(true, sli)
	// a := plus(true, iMat, z)
	// b := getReverse(a)
	// c := plus(true, b, iMat)
	// return c
	// sli := make([]float64, 6)
	// for i := 0; i < 6; i++ {
	// 	for j := 0; j < 6; j++ {
	// 		sli[i] += readFloat()
	// 	}
	// }
	// return sli
	return exponent()
}

func exponent() interface{} {
	fmt.Println("read int to exponent")
	m := readInt()
	// s := read()

	base := make([][]int, n)
	sli := make([][]int, n)
	fmt.Println("read matrix to exponent")
	for i := 0; i < n; i++ {
		base[i] = readSli(n)
		sli[i] = make([]int, n)
		copy(sli[i], base[i])
	}
	var calculate = func(left, right [][]int, row, col int) int {
		sum := 0
		for i := 0; i < n; i++ {
			sum += left[row][i] * right[i][col]
		}
		return sum
	}

	var ans [][]int
	for i := 1; i < m; i++ {
		ans = make([][]int, n)
		for j := 0; j < n; j++ {
			ans[j] = make([]int, n)
			for k := 0; k < n; k++ {
				ans[j][k] = calculate(sli, base, j, k)
			}
		}
		copy(sli, ans)
	}

	return ans
}

func multiply(div bool, arg1 interface{}) interface{} {
	fmt.Println("read float to multiply")
	m := readFloat()

	if div {
		m = 1 / m
	}
	var sli [][]float64
	if arg1, ok := arg1.([][]float64); ok {
		sli = arg1
	} else {
		fmt.Println("read matrix to multiply")
		sli = readMatrix()
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sli[i][j] *= m
		}
	}
	return sli
}

func plus(minus bool, arg1, arg2 interface{}) interface{} {
	sli1 := make([][]float64, n)
	sli2 := make([][]float64, n)
	if arg1, ok := arg1.([][]float64); ok {

		sli1 = arg1
	} else {
		fmt.Println("read matrix to plus")
		sli1 = readMatrix()
	}
	if arg2, ok := arg2.([][]float64); ok {
		sli2 = arg2
	} else {
		fmt.Println("read matrix to plus")
		sli2 = readMatrix()
	}

	var calc func(a, b float64) float64
	if minus {
		calc = func(a, b float64) float64 {
			return a - b
		}
	} else {
		calc = func(a, b float64) float64 {
			return a + b
		}
	}

	ans := make([][]float64, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			ans[i][j] = calc(sli1[i][j], sli2[i][j])
		}
	}
	return ans
}

func getReverse(arg interface{}) [][]float64 {
	var sli [][]float64
	if arg, ok := arg.([][]float64); ok {
		sli = arg
	} else {
		fmt.Println("read matrix to reverse")
		sli = readMatrix()
	}

	rev := getIdentityMatrix(n)
	// execute n times
	for i := 0; i < n; i++ {
		// swap for all rows, and save base row
		j := 0
		for j = i; j < n; j++ {
			if sli[j][i] == 0 {
				continue
			}

			sli[i], sli[j] = sli[j], sli[i]
			rev[i], rev[j] = rev[j], rev[i]
			break
		}
		// change target to 1
		coTo1 := sli[i][j]
		for k := 0; k < n; k++ {
			sli[i][k] /= coTo1
			rev[i][k] /= coTo1
		}

		// add for all rows
		for k := 0; k < n; k++ {
			// no add base row
			if k == j {
				continue
			}

			co := sli[k][i] * -1
			// for all columns
			for l := 0; l < n; l++ {
				sli[k][l] += co * sli[j][l]
				rev[k][l] += co * rev[j][l]
			}
		}
	}
	return rev
}

func getIdentityMatrix(n int) [][]float64 {
	sli := make([][]float64, n)
	for i := 0; i < n; i++ {
		sli[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			if i == j {
				sli[i][j] = 1
			}
		}
	}
	return sli
}

// ========================read
// func read() string {
// 	sc.Scan()
//     return sc.Text()
// }

func readMatrix() [][]float64 {
	sli := make([][]float64, n)
	for i := 0; i < n; i++ {
		sli[i] = readFloatSli(n)
	}
	return sli
}

func readFloat() float64 {
	sc.Scan()
	ret, _ := strconv.ParseFloat(sc.Text(), 64)
	return ret
}

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

func readFloatSli(n int) []float64 {
	// sli := make([]string, n)
	sli := make([]float64, n)
	for i := 0; i < n; i++ {
		// sli[i] = read()
		sli[i] = readFloat()
	}
	return sli
}

func readInt() int {
	sc.Scan()
	ret, _ := strconv.Atoi(sc.Text())
	return ret
}

// =======================main========================
func main() {
	buf := make([]byte, initial_buf)
	sc.Buffer(buf, max_bufSize)
	sc.Split(bufio.ScanWords)
	result := run()
	// print(result)
	p(result)
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
		fmt.Println(v)
	}
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
