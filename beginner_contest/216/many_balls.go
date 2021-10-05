package main

import (
	"fmt"
	// "bufio"
	// "os"
	// "strings"
	// "strconv"
	// "reflect"
	// "math"
	// "sort"
)

// var sc = bufio.NewScanner(os.Stdin)

func main() {

	var n int64
	fmt.Scan(&n)
	ans := ""
	for i := 0; i < 120; i++ {
		// fmt.Println(ans, n)
		if n % 2 == 1 {
			ans = "A" + ans
			if n == 1 {
				fmt.Println(ans)
				return
			}
			n = n - 1
		} else {
			n = n / 2
			ans = "B" + ans
		}
	}
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

// func nextInt() int {
//     sc.Scan()
//     ret, _ := strconv.Atoi(sc.Text())
//     return ret
// }

// func (a SortBy) Len() int           { return len(a) }
// func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortBy) Less(i, j int) bool { return a[i].t < a[j].t }