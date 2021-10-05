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
	var a, b, c, x, sum int
	fmt.Scan(&a, &b, &c, &x)
	for i := 0; i <= a; i++ {
		// fmt.Println("i:", i)
		for j := 0; j <= b; j++ {
			// fmt.Println("j:", j)
			for k := 0; k <= c; k++ {
				// fmt.Println("k:", k)
				if (i * 500 + j * 100 + k * 50 == x) {
					sum += 1
				}
			}
		}
	}
	// fmt.Println(x, a, b, c)
	fmt.Println(sum)
}