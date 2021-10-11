package math

func Abs(x, y int) float64 {
	// return absolute value
}

func Pow(x, y int) float64 {
	// return x^y
}


package original
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

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}