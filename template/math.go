package custom_math

import (
	"math"
	"sort"
)

func Abs(x float64) float64 {
	return math.Abs(x)
}

func DiffAbs(x, y int) int {
	return int(math.Abs(float64(x - y)))
}

func Pow(x, y int) float64 {
	return math.Pow(float64(x), float64(y))
}

func RoundUp(x, y int) int {
	return (x + y - 1) / y
}

func SumXtoY(x, y int) int {
	sum := x + y
	if sum%2 == 0 {
		return sum*2 + sum/2
	}
	return sum * 2
}

func Sum(sli []int) int {
	sum := 0
	for _, i := range sli {
		sum += i
	}
	return sum
}

func MinSli(sli []int) int {
	sort.Ints(sli)
	return sli[0]
}

func MaxWithIdx(x, y, ix, iy int) (int, int) {
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

// factorization
var memo map[int]int

func factorization(n int) {
	prime_numbers := []int{2, 3}
	memo[1] = 1
	memo[2] = 2
	memo[3] = 2
	for i := 4; i <= n; i++ {
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
		} else {
			ret := 1
			for _, c := range m {
				ret *= (c + 1)
			}
			memo[i] = ret
		}
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
