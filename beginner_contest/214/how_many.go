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
	var s, t int 
	fmt.Scan(&s, &t)

	combi := 0
	for i := 0; i <= s; i++ {
		for j := 0; i + j <= s; j++ {
			for k := 0; i + j + k <= s; k++ {
				if i * j * k <= t {
					combi += 1
				}
			}
		}
	}
	fmt.Println(combi)
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