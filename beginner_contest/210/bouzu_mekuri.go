package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
	// "strconv"
	// "reflect"
	// "math"
	// "sort"
)

// var sc = bufio.NewScanner(os.Stdin)

func main() {
	var n int
	var s string
	fmt.Scan(&n, &s)
	// sc.Split(bufio.ScanWords)
	for pos, r := range s {
		if r == '1' {
			if pos % 2 == 0 {
				fmt.Println("Takahashi")
				return 
			} else {
				fmt.Println("Aoki")
				return
			}
		}
	}
}


// func next() string {
// 	sc.Scan()
//     ret := sc.Text()
//     return ret
// }

// func nextInt() int {
//     sc.Scan()
//     ret, _ := strconv.Atoi(sc.Text())
//     return ret
// }