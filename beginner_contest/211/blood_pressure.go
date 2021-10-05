package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
	// "strconv"
	// "reflect"
)

func main() {
	var a, b float64 
	fmt.Scan(&a, &b)
	var c float64
	c = (a - b) / 3.0 + b
	fmt.Printf("%g", c)
}