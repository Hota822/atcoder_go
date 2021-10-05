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
	var x string
	fmt.Scan(&x)
	// sc.Split(bufio.ScanWords)
	if x == "Hello,World!" {
		fmt.Println("AC")
		return
	}
	fmt.Println("WA")
	
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
9223372036854775807