package main

import (
	"fmt"
)

func main() {
	var x uint
	fmt.Scan(&x)
	if x > 110 {
		fmt.Println(3)
		return
	}
	if x == 100 || x == 1 || x == 8 {
		fmt.Println(1)
		return
	}
	if x == 0 {
		fmt.Println(0)
		return
	}
	fmt.Println(2)
}