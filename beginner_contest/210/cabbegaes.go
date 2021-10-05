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
	var n, a, x, y int64
	fmt.Scan(&n, &a, &x, &y)
	// sc.Split(bufio.ScanWords)
	if n > a {
		fmt.Println((x * a) + (n-a) * y)
	} else {
		fmt.Println(x * n)
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