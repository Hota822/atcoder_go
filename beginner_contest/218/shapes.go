package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"reflect"
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

	s_sli := make([][]int, n)
	s_rng := [4]int {0, 0 , 0, 0}
	// left up y, x, right down y, x

	t_sli := make([][]int, n)
	t_rng := [4]int {0, 0 , 0, 0}

	for i := 0; i < n; i++ {
		s_sli[i] = make([]int, n)
		str := next()
		idx := strings.Index(str, "#")
		fmt.Println(idx, str)
		if idx >= 0 {
			if s_rng[0] == 0 {
				s_rng[0] = i
			}
			if s_rng[1] > idx {
				s_rng[1] = idx
			}
			if s_rng[2] < i {
				s_rng[2] = i
			}
			if s_rng[3] < idx {
				s_rng[3] = idx
			}
		}
		for pos, j := range str {
			if j == '.' {
				s_sli[i][pos] = 0
			} else {
				s_sli[i][pos] = 1
			}
		}
	}
	
	for i := 0; i < n; i++ {
		t_sli[i] = make([]int, n)
		str := next()
		idx := strings.Index(str, "#")
		if idx >= 0 {
			if t_rng[0] == 0 {
				t_rng[0] = i
			}
			if t_rng[1] > idx {
				t_rng[1] = idx
			}
			if t_rng[2] < i {
				t_rng[2] = i
			}
			if t_rng[3] < idx {
				t_rng[3] = idx
			}
		}
		for pos, j := range str {
			if j == '.' {
				t_sli[i][pos] = 0
			} else {
				t_sli[i][pos] = 1
			}
		}
	}
	fmt.Println(s_sli)
	fmt.Println(t_sli)
	fmt.Println(s_rng, t_rng)
	s_sli = s_sli[s_rng[1]:s_rng[3] + 1]
	t_sli = t_sli[t_rng[1]:t_rng[3] + 1]
	fmt.Println(s_sli)
	fmt.Println(t_sli)
	s_ans := make([][]int, n)
	t_ans := make([][]int, n)
	for i:= 0; i < len(s_sli); i ++ {
		s_ans[i] = s_sli[i][s_rng[0]:s_rng[2] + 1]
		t_ans[i] = t_sli[i][t_rng[0]:t_rng[2] + 1]
	}

	// fmt.Println(s_ans)
	// fmt.Println(t_ans)

	// for i := 0; i < n; i++ {
	// 	t_sli[i] = make([]int, n)
	// 	for j := 0; j < n; j++ {
	// 		t_sli[i][j] = next()
	// 	}
	// }

	if reflect.DeepEqual(s_ans, t_ans) {
		fmt.Println("Yes")
		return
	}

	t90_sli := spin90(t_ans, n)
	if reflect.DeepEqual(s_ans, t90_sli) {
		fmt.Println("Yes")
		return
	}

	t180_sli := spin180(t_ans, n)
	if reflect.DeepEqual(s_ans, t180_sli) {
		fmt.Println("Yes")
		return
	}

	t270_sli := spin180(t90_sli, n)
	if reflect.DeepEqual(s_ans, t270_sli) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

// func swap(sli *[]string, i, j int) {
// 	a, b := (*sli)[i], (*sli)[j]
// 	(*sli)[i], (*sli)[j] = b, a
// }


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

// func (a SortBy) Len() int           { return len(a) }
// func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a SortBy) Less(i, j int) bool { return a[i].t < a[j].
func spin90(sli [][]int, n int) [][]int {
	ret := make([][]int, n)
	for i := 0; i < n; i ++ {
		ret[i] = make([]int, n)
	}
	for pos_i, i := range sli {
		for pos_j, j := range i {
			ret[pos_j][pos_i] = j
		}
	}
	return ret
}

func spin180(sli [][]int, n int) [][]int {
	len := len(sli)
	ret := make([][]int, n)
	for i := 0; i < len; i++ {
		ret[i] = sli[len - 1 - i]
	}
	return ret
}

// func moveVer(sli [][]int, n int) [][]int {

// }

// func moveHor(sli [][]int, n int) [][]int {

// }