package math

// math package
func Abs(x float64) float64 {
	// return absolute value
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return - x
}

func AbsInt(x int) int {
	return int(math.Abs(float64(x)))
}

func DiffAbs(x, y int) int {
	return int(math.Abs(float64(x - y)))
}

// math package
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

func MaxWithIdx(x, y ,ix, iy, int) (int, int) {
	if x > y {
		return x, ix
	}
	return y, iy
}

func MaxSli(sli []int) int {
	sort.Ints(sli)
	l := len(sli)
	return sli[l]
}

func MaxOf3(sli [3]int) int {
	x, y, z := sli[0], sli[1], sli[2]
	if x > y {
		if x > z {
			return x
		}
	} else {
		if y > z {
			return y
		}
	}
	return z
}