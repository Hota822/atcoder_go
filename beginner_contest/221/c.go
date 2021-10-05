package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	"strconv"
	// "reflect"
	// "math"
	"sort"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	n := next()
	sli := make([]int, len(n))
	var a, b string
	// var c, d string
	for pos, c := range n {
		sli[pos] = convInt(string(c))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sli)))
	var last string
	ki := (len(n) % 2 == 1)
	// fmt.Println(ki)
	for pos, i := range sli {
		s := convStr(i)
		if pos % 2 == 0 {
			if ki && len(n) == pos + 1 {
				last = s
			} else {
				a += s
			}
			// d += s
		} else {
			
			// } else {
				b += s
				// c += s
			// }
		}
	}
	if ki {
		// ans_ab1 := convInt(a + last) * convInt(b)
		// ans_ab2 := convInt(a) * convInt(b + last)
		// ans_cd1 := convInt(c + last) * convInt(d)
		// ans_cd2 := convInt(c) * convInt(d + last)
		// ans_sli := []int {ans_ab1, ans_ab2, ans_cd1, ans_cd2}
		ans_x := convInt(a) * convInt(b + last)
		ans_y := convInt(a + last) * convInt(b)
		// sort.Sort(sort.Reverse(sort.IntSlice(ans_sli)))
		if ans_x > ans_y {
			fmt.Println(ans_x)
		} else {
			fmt.Println(ans_y)
		}
		// fmt.Println(ans_sli[0])
	} else {
		ans_ab := convInt(a) * convInt(b)
		c, d := swap(a, b)
		ans_cd := convInt(c) * convInt(d)
		// fmt.Println(a, b, c, d)
		if ans_ab > ans_cd {
			fmt.Println(ans_ab)
		} else {
			fmt.Println(ans_cd)
		}
	}
	// fmt.Println(len(n))
}

func next() string {
	sc.Scan()
    ret := sc.Text()
	return ret
}

func convInt(s string) int {
    ret, _ := strconv.Atoi(s)
    return ret
}

func convStr(i int) string {
	return strconv.Itoa(i)
}

func swap(a, b string) (string, string) {
	a_len := len(a)
	b_len := len(b)
	la := a[a_len -1]
	lb := b[b_len -1]
	return a[:(a_len -1)] + string(lb), b[:(b_len -1)] + string(la)
}