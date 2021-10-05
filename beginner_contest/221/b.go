package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	// "strconv"
	// "reflect"
	// "math"
	// "sort"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

	s, t := next(), next()
	if s == t {
		fmt.Println("Yes")
		return
	}
	len := len(s)
	ope := false
	after_change := false
	for pos, c := range s {
		if byte(c) != t[pos] && !after_change {
			if pos == len {
				fmt.Println("No")
				return
			}

			// 0 times
			if !ope {
				ope = true
				// changed
				if byte(c) == t[pos +1] && s[pos +1] == t[pos] {
					after_change = true
					continue
				} else {
					fmt.Println("No")
					return
				}
			} else {
				fmt.Println("No")
				return
			}
		}
		after_change = false
	}
	fmt.Println("Yes")
	// sli := make([]string, 3)
	// for i := 0; i < 3; i++ {
	// 	sli[i] = next()
	// }
	// t := next()
	// for _, i := range t {
	// 	idx, _ := strconv.Atoi(string(i))
	// 	fmt.Print(sli[idx -1])
	// }
	// fmt.Println("")
}


func next() string {
	sc.Scan()
    ret := sc.Text()
    return ret
}

// func nextInt() int {
//     sc.Scan()
//     ret, _ := strconv.Atoi(sc.Text())
//     return ret
// }