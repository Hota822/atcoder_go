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
	fmt.Scan(&n)
	if n >= 212 {
		fmt.Println(8)
		return 
	} else if n >= 126 {
		fmt.Println(6)
		return 
	}
	fmt.Println(4)
	// sc.Split(bufio.ScanWords)
	
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