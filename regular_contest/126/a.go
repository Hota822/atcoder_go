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
	t := nextInt()
	sli := make([][3]int64, t)
	for i:=0; i < t; i++ {
		sli[i] = [3]int64 {nextInt64(), nextInt64(), nextInt64()}
	}
	for _, v := range sli {
		// fmt.Println("a")
		fmt.Println(count(v[0], v[1], v[2]))
	}
}

func count(n2, n3, n4 int64) int64 {
	n3p2 := n3 / 2
	n2p2 := n2 / 2
	n2mod := n2 % 2
	// if n3a > n4 + n2a {
	// 	return n3a - n4 - n2a
	// }
	// if n3a - n4
	if n3p2 > n4 {
		n3r := n3p2 - n4
		n10_1st := n4
		if n3r > n2p2 {
			n10_2nd := n2p2
			// return n10 + n3r - n2p2
			return n10_1st + n10_2nd
		} else {
			n10 := n3p2
			// n10_2nd :=
			n2r := n2p2 - n3r
			n2r_total := n2r * 2 + n2mod
			return n10 + n2r_total / 5
		}
	} else {
		n10_1st := n3p2
		n4r := n4 - n3p2
		n4rp2 := n4r / 2
		n4r_mod := n4r % 2
		if n4rp2 > n2 {
			n10_2nd := n2
			return n10_1st + n10_2nd
		} else {
			n10_2nd := n4rp2
			n2r := n2 - n4rp2
			mod := n4r_mod * 2 + n2r
			n10_3rd := mod / 5
			return n10_1st + n10_2nd + n10_3rd
		}
	}

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

func nextInt64() int64 {
	var ret int64
    sc.Scan()
    ret, _ = strconv.ParseInt(sc.Text(), 10, 64)
    return ret
}

