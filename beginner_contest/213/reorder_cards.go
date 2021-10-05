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
	num   int
	coords int
}

type SortBy []card

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].coords < a[j].coords }

func readInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func run(r io.Reader) string {
	sc = bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)

	// read h, w, n
	_, _, n := readInt(), readInt(), readInt()

	// make a_slice, b_slice
	a_sli, b_sli := make([]card, n), make([]card, n)

	// read ai, bi
	for i := range a_sli {
		a_sli[i] = card{i: i, val: readInt()}
		b_sli[i] = card{i: i, val: readInt()}
	}

	// sort by ai or bi
	sort.Sort(SortBy(as))
	sort.Sort(SortBy(bs))

	// output value container
	ans_a := make([]int, n)
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