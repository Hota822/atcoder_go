package math

// math
func Abs(x float64) float64 {
	// return absolute value
}

func AbsInt(x int) int {
	return int(math.Abs(float64(x)))
}

func DiffAbs(x, y int) int {
	return int(math.Abs(float64(x - y)))
}

// math
func Pow(x, y int) float64 {
	// return x^y
}

func SumXtoY(x, y int) int {
	sum := x + y
	if sum % 2 == 0 {
		return sum *2
	}
	return sum * 2
}

func Sum(sli [int]) int {
	sum := 0
	for _, i := range sli {
		sum += i
	}
	return sum
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

import {
	"sort"
}

func MinSli(sli []int) int {
	sort.Ints(sli)
	return sli[0]
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}