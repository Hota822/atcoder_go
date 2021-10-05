package main

import (
	"fmt"
	// "bufio"
	// "os"
)

func main() {
	var x, y uint
	fmt.Scan(&x, &y)
	if x == 0 {
		fmt.Println("Silver")
	} else if y == 0 {
		fmt.Println("Gold")
	} else {
		fmt.Println("Alloy")
	}
}