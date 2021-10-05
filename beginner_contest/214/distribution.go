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

var sc = bufio.NewScanner(os.Stdin)

type Ts struct {
	i int
	t int64
}

type SortBy []Ts

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].t < a[j].t }


func main() {
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	s_sli := make([]int64, n)

	// t_sli = make([]int64, n)
	sort_by := make([]Ts, n)
	for i := 0; i < n; i++ {
		s_sli[i] = nextInt()
	}
	for i := 0; i < n; i++ {
		sort_by[i] = Ts { i: i, t: nextInt()}
	}

	sort.Sort(SortBy(sort_by))

	// var now int64 = sort_by[0].t
	// var now int64 = 0
	ans := make([]int64, n + 1)
	ans[sort_by[0].i] = sort_by[0].t
	for i := 1; i < n - 1; i++ {
		t := &sort_by[i]
		if ans[t.i] > t.t {
			ans[t.i] = t.t
		}
		// now = t.t
		next := t.i + 1
		if t.i + 1 == n {
			next = 1
		}
		ans[next] = ans[t.i] + s_sli[t.i]
	}
	for i := 0; i < n; i++ {
		fmt.Println(ans[i])
	}
}

func nextInt() int64 {
	var ret int64
    sc.Scan()
    ret, _ = strconv.ParseInt(sc.Text(), 10, 64)
    return ret
}