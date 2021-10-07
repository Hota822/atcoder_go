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
	str := []string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	var p int
	sc.Split(bufio.ScanWords)
	for i := 0; i < 26; i++ {
		fmt.Scan(&p)
		fmt.Print(str[p -1])
	}
	fmt.Println("")
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