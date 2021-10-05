package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
	// "strconv"
	// "reflect"
	// "math"
)

func main() {
	var s string
	fmt.Scan(&s)
	// sum := 0
	count := [9] uint64{1, 0, 0, 0, 0, 0, 0, 0, 0, }
	for _, chr := range s {
		for pos, t := range "chokudai" {
			if t == chr {
				count[pos + 1] += count[pos]
				count[pos + 1] %= 1000000007
				break
			}
		}
	}
	fmt.Println(count[8])
}
