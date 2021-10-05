package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	"strconv"
	// "reflect"
	// "math"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	// var x int 
	// fmt.Scan(&x)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	last := [2]int {0, 0}
	boobee := [2]int {0, 0}
	for i := 0; i < n; i++ {
		j := nextInt()
		if j > boobee[0] {
			if j > last[0] {
				boobee[0] = last[0]
				boobee[1] = last[1]
				last[0] = j
				last[1] = i
			} else {
				boobee[0] = j
				boobee[1] = i
			}
		}
	}
	fmt.Println(boobee[1] + 1)
}


// func next() string {
// 	sc.Scan()
//     ret := sc.Text()
//     return ret
// }

func nextInt() int {
    sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
    return ret
}