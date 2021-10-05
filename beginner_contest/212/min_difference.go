package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
	"math"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n, m, result int
	n = nextInt()
	_ = nextInt()
	var is, js string
	is = sc.Text()
	js = sc.Text()
	min := math.Pow(10, 9)

	for ipos, ic := range is {
		for _, jc := range js {
			result = math.Abs(strconv.Atoi(string([]rune{ic})) - strconv.Atoi(string([]rune{jc})))
			if min > result {
				min = result
			}
		}
    }

	fmt.Printf(result)
}

func nextInt() int {
    sc.Scan()
    i, e := strconv.Atoi(sc.Text())
    return i
}