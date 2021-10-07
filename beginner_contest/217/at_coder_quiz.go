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
	sli := make([]string, 4)

	for i:=0; i < 3; i++ {
		s := next()
		if s == "ABC" {
			sli[0] = s
		}
		if s == "ARC" {
			sli[1] = s
		}
		if s == "AGC" {
			sli[2] = s
		}
		if s == "AHC" {
			sli[3] = s
		}
	}
	for j := 0; j < 4; j ++ {
		if sli[j] == "" {
			if j == 0 {
				fmt.Println("ABC")
			}
			if j == 1 {
				fmt.Println("ARC")
			}
			if j == 2 {
				fmt.Println("AGC")
			}
			if j == 3 {
				fmt.Println("AHC")
			}
		}
	}
}


func next() string {
	sc.Scan()
    ret := sc.Text()
    return ret
}