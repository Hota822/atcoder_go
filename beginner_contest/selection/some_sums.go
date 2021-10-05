package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
	// "strconv"
	// "reflect"
)

func main() {
	var n, a, b, sum int
	fmt.Scan(&n, &a, &b)
	// n_len := len(strconv.Itoa(n))
	// for j := 0, j < n / 10; j++ {
	// 	for i := a + j * 9, i <= b + j * 9; i++ {
	// 		sum += i
	// 	}
	// }
	// fmt.Println(n_len, a, b)
	for {
		for i := 1; i <= n; i++ {
			if i >= a {
				if n <= i {
					fmt.Println(sum)
				}
				if i > b {
					a += 9
					b += 9
					break
				} else {
					sum += i
				}
			}
		}
	}
}