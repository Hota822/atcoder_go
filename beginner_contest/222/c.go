package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	// "reflect"
	// "math"
	"sort"
)

const (
    // maxBufSize = 100000
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	// buf := make([]byte, maxBufSize)
	// sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords)

	n, m := readInt(), readInt()
	sli := make([]Man, 2 * n)

	for i:=0; i < 2 * n; i++ {
		sli[i] = Man {n: i +1, order: 0, te: strings.Split(read(), "")}
		// fmt.Println(sli[i].te)
	}

	for i:=0; i < m; i++ {
		for j:=0; j < 2 * n; j += 2 {
			a := &sli[j]
			b := &sli[j +1]
			win := Win(a.te[i], b.te[i])
			if win == 1 {
				a.order--
			} else if win == -1 {
				b.order--
			}
		}
		sort.Sort(SortBy(sli))
	}
	// y := readInt()

	// s := read()
	// t := read()

	// n := readInt()
	// sli := make([]int, n)
	// sli := make([]string, n)

	return sli
}

func Win(a, b string) int {
	if a == "G" {
		if b == "C" {
			return 1
		} else if b == "P" {
			return -1
		}
		return 0
	}
	if a == "C" {
		if b == "P" {
			return 1
		} else if b == "G" {
			return -1
		}
		return 0
	}
	if a == "P" {
		if b == "G" {
			return 1
		} else if b == "C" {
			return -1
		}
		return 0
	}
	return 0
}

type SortBy []Man

type Man struct {
	n int
	te []string
	order int
}

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool {

	if a[i].order == a[j].order {
		return a[i].n < a[j].n
	}
	return a[i].order < a[j].order
}





// ========================read
func read() string {
	sc.Scan()
    ret := sc.Text()
    return ret
}

func readInt() int {
    sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
    return ret
}


//=======================main========================
func main() {
	result := run()
	print(result)
}

func print(ans interface{}) {
	if v, ok := ans.(bool); ok {
		if v {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		return
	}
	if sli, ok := ans.([]int); ok {
		for _, v := range sli {
			fmt.Println(v)
		}
		return
	}
	if sli, ok := ans.([]int64); ok {
		for _, v := range sli {
			fmt.Println(v)
		}
		return
	}
	if sli, ok := ans.([]Man); ok {
		for _, v := range sli {
			fmt.Println(v.n)
		}
		return
	}
	fmt.Println(ans)
}
