package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
	"strconv"
	// "reflect"
	// "math"
	// "sort"
)

// var sc = bufio.NewScanner(os.Stdin)

func main() {
	// var n float64 
	fmt.Scan(&n)
	// sc.Split(bufio.ScanWords)
	fmt.Println(int(math.Log2(n)))
	a := 59.9999999999999999999999999999999999999
	// fmt.Println(strconv.FormatFloat(a,'f',  -1, 64))
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

import math
from sys import stdin

n = stdin.readline().split()
print(math.log2(n))