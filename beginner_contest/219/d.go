package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sc *bufio.Scanner

type card struct {
	x int
	y int
}

type SortBy []card

func (a SortByY) Len() int           { return len(a) }
func (a SortByY) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByY) Less(i, j int) bool { return a[i].y < a[j].y }

func (a SortByX) Len() int           { return len(a) }
func (a SortByX) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByX) Less(i, j int) bool { return a[i].x < a[j].x }

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	// read h, w, n
	n, x, y := readInt(), readInt(), readInt()

	// make a_slice, b_slice
	a_sli, b_sli := make([]card, n), make([]card, n)

	// read ai, bi
	for i := 0; i < n; i++ {
		ax, ay := readInt(), readInt()
		c := card{x: ax, y: ay}
		a_sli[i] = c
		b_sli[i] = c
	}

	// sort by ai or bi
	sort.Sort(SortByX(a_sli))
	sort.Sort(SortByY(b_sli))

	// output value container
	var ans1 [2]int
	var ans2 [2]int
	for i := 0; i < n; i++ {
		if ans1[1] > x && ans1[2] > y {
			return i
		}
		if ans2[1] > x && ans2[2] > y {
			return i
		}
		ans1[0] += a_sli[i].x
		ans2[1] += 
	}
	
	ans_b := make([]int, n)

	// banpei
	idx_a, idx_b := 1, 1
	for i := 0; i < n; i++ {
		// a card { num, coords }
		a, b := a_sli[i], b_sli[i]

		// not first time process
		if i > 0 {
			// if current is different form before, add index
			if a_sli[i-1].coords != a.coords {
				idx_a++
			}
			if b_sli[i-1].coords != b.coords {
				idx_b++
			}
		}

		// reorder first index by result
		ans_a[a.num] = idx_a
		ans_b[b.num] = idx_b
	}

	// make output slice
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		ans[i] = fmt.Sprintf("%v %v", ans_a[i], ans_b[i])
	}
	return strings.Join(ans, "\n")
}