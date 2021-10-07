package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	"strconv"
	// "reflect"
	// "math"
	// "sort"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)

	sli := make([]string, 3)
	for i := 0; i < 3; i++ {
		sli[i] = next()
	}
	t := next()
	for _, i := range t {
		idx, _ := strconv.Atoi(string(i))
		fmt.Print(sli[idx -1])
	}
	fmt.Println("")
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