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
	for pos, i := range s {
		if pos + 1 == n {
			if string(i) == "o" {
				fmt.Println("Yes")
				return
			}
			fmt.Println("No")
		}
	}

	// if s < t {
	
	// 	return 
	// }
	// fmt.Println(s[2]);
	// // sc.Split(bufio.ScanWords)
	// sli := strings.Split(s, ".")
	// x := sli[0]
	// y, _ := strconv.Atoi(sli[1])
	// if y <= 2 {
	// 	fmt.Println(x + "-")
	// 	return
	// }
	// if y <= 6 {
	// 	fmt.Println(x)
	// 	return
	// }
	// fmt.Println(x + "+")
	// fmt.Println("WA")
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
