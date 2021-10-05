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
const (
    // この値がinitialBufSize以下の場合、Scannerはバッファの拡張を一切行わず与えられた初期バッファのみを使う。
    maxBufSize = 100000
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, maxBufSize)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords)
	n := nextInt()
	sli := make([]string, n)
	// first := 0
	for i := 0; i < n; i++ {
		p := nextInt()
		// if first == p - 1 {
		// 	fmt.Print(strconv.Itoa(i + 1) + " ")
		// 	// fmt.Println("first")
		// 	first++
		// } else {
			// fmt.Println("else")
			sli[p - 1] = strconv.Itoa(i + 1)
		// }
	}
	for _, s := range sli[:n - 1] {
		fmt.Print(s)
		fmt.Print(" ")
	}
	fmt.Print(sli[n - 1])
	// len := len(ans)
	// if len != 0 {
	// 	// fmt.Print(ans)
	// 	fmt.Println(ans[:len - 1])
	// }
	fmt.Print("\n")

	// before
	// ans := ""
	// ans_sli := sli[first:n]
	// for _, s := range ans_sli {
	// 	with_space := s + " "
	// 	ans += with_space
	// }
	// len := len(ans)
	// if len != 0 {
	// 	// fmt.Print(ans)
	// 	fmt.Println(ans[:len - 1])
	// }
	// fmt.Print("\n")
}

// func swap(sli *[]string, i, j int) {
// 	a, b := (*sli)[i], (*sli)[j]
// 	(*sli)[i], (*sli)[j] = b, a
// }


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

// func (a SortBy) Len() int           { return len(a) }
// func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortBy) Less(i, j int) bool { return a[i].t < a[j].t }