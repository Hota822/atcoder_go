package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
	// "strconv"
	// "reflect"
	// "math"
)

// var sc = bufio.NewScanner(os.Stdin)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := 0; i < 256; i ++ {
		if a ^ i == b {
			fmt.Println(i)
		}
	}
	// fmt.Println(x)
}