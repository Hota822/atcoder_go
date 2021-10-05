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
	var x int
	fmt.Scan(&x)
	if x >= 90 {
		fmt.Println("expert")
		return;
	}
	if x >= 70 {
		dif(x, 90)
		return
	}
	if x >= 40 {
		dif(x, 70)
		return
	}
	dif(x, 40)
}

func dif(x, y int) {
	ans := y - x
	fmt.Println(ans)
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
