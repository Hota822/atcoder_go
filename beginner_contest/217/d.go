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
	l, q := nextInt64(), nextInt64()
	sli := make([]int, l)
	var i int64;
	for i = 0 ; i < q; i++ {
		c := nextInt()
		x := nextInt64()
		if c == 1 {
			sli[i] = 1
		} else {
			var j int64;
			len := 0
			for j = 1; x + j < l;j++ {
				if sli[x + j] == 0{
					len++
				} else {
					break
				}
			}
			for j = 1; x - j > 0 ;j++ {
				if sli[x - j] == 0{
					len++
				} else {
					break
				}
			}
			fmt.Println(len)
		}
	}
}

// type Pair []Ball

// type Ball struct {
// 	idx int
// 	tube int
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

func nextInt64() int64 {
	var ret int64
    sc.Scan()
    ret, _ = strconv.ParseInt(sc.Text(), 10, 64)
    return ret
}