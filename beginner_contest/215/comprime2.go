package main

import (
	"fmt"
	// "bufio"
	// "os"
	"strings"
	// "strconv"
	// "reflect"
	// "math"
	"sort"
)

// var sc = bufio.NewScanner(os.Stdin)

func main() {
	var s string
	fmt.Scan(&s)
	var k int
	if len(s) == 1 {
		fmt.Println(s)
	}
	fmt.Scan(&k)
	// sc.Split(bufio.ScanWords)
	sli := strings.Split(s, "")
	l := len(sli)
	// fmt.Println(sli, s, k)
	sort.Strings(sli)
	// last := make([]string, 0, 8)
	// ty := 1
	// be := 1
	times := 1
	ans_sli := make([]string, 0, 8)
	ans := ""
	for i := 0; i < k; i ++{
		ans_lis = copy(sli)
		if times == k {
			break
		}
		swap(ans_sli, 0, i)
		for j := 0; j < k - 1[ j++ {
			if times == k {
				break
			}
			swap(ans_sli, 1, i + j)
			for m := 0; m < k - 2[ m++ {
				if times == k {
					break
				}
				swap(ans_sli, 2, i + j + m)
				for n := 0; n < k - 2[ n++ {
					if times == k {
						break
					}
					swap(ans_sli, 3, i + j + n)
					for o := 0;o < k - 2[ o++ {
						if times == k {
							break
						}
						swap(ans_sli, 4, i + j + n)
					}
				}
			}
		}
		// t_1 := sli[(l-2):(l-1)]
		times++
		swap(&ans_sli, l-2, l-1)
		fmt.Println(sli)
		// fmt.Println(t_1)
	}
	
	for _, a := range sli {
		ans += a
	}
	// fmt.Println(last)
	fmt.Println(ans)
	
}

func swap(sli *[]string, i, j int) {
	a, b := (*sli)[i], (*sli)[j]
	(*sli)[i], (*sli)[j] = b, a
}


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