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
	var a, b int
	fmt.Scan(&a, &b)
	diff := a - b
	if diff == 0 {
		fmt.Println(1)
		return
	}
	var e int64 = 1
	for i := 0; i < diff; i++ {
		e *= 32
	}
	fmt.Println(e)
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
