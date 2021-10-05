package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	"strconv"
	// "reflect"
	// "math"
	"sort"
)
const (
    // この値がinitialBufSize以下の場合、Scannerはバッファの拡張を一切行わず与えられた初期バッファのみを使う。
    maxBufSize = 500000
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	buf := make([]byte, maxBufSize)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	sli := make([][2]int, m)

	for i := 0; i < m; i++ {
		sli[i] = [2]int {nextInt(), nextInt()}
	}
	for 
	x_map := make(map[rune]int, 26)
	for pos, r := range x {
		x_map[r] = pos + 1
	}
	// fmt.Println(x_map)
	sli := make([]name, n)

	for i := 0; i < n; i++ {
		s := next()
		var name_sli [10]int
		// name_sli := [10]int {26, 26, 26, 26, 26, 26, 26, 26, 26, 26}
		for pos, j := range s {
			name_sli[pos] = x_map[j]
		}
		sli[i] = name{str: s, sli: name_sli}
		//  = card{i: i, val: readInt()}
	}
	// for _, item := range sli {
	// 	fmt.Println(item.sli)
	// }

	sort.Sort(SortBy(sli))
	for _, item := range sli {
		fmt.Println(item.str)
	}
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

type SortBy []name

type name struct {
	str string
	sli [10]int
}

func (nm SortBy) Len() int           { return len(nm) }
func (nm SortBy) Swap(i, j int)      { nm[i], nm[j] = nm[j], nm[i] }
func (nm SortBy) Less(i, j int) bool {
	for a:= 0; a < 10; a ++ {
		if nm[i].sli[a] == nm[j].sli[a] {
			continue
		}
		return nm[i].sli[a] < nm[j].sli[a]
	}
	return false
}