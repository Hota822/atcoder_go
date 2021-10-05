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
	n := nextInt()
	sli := make([][2]string, n)
	for i:= 0; i < n; i++ {
		sei := next()
		mei := next()
		for _, val := range sli {
			// fmt.Println(val[0], val[1])
			if val[0] == sei && val[1] == mei {
				fmt.Println("Yes")
				return
			}
		}
		sli[i] = [2]string {sei, mei}
	}

	fmt.Println("No")
}


func next() string {
	sc.Scan()
    ret := sc.Text()
    return ret
}

func nextInt() int {
    sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
    return ret
}