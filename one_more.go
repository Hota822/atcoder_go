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
var sli_len int

func main() {
	var s string
	var k int
	fmt.Scan(&s, &k)
	sli_len = len(s)
	all := GetCombination(s)
	fmt.Println(all[k -1])
}

func GetPermutation(s string) []string {
	swap := func(i, j int, sli []string) []string {
		new := make([]string, len(sli))
		copy(new, sli)
		new[i], new[j] = sli[j], sli[i]
		return new
	}

	var n int = 1
	l := len(s)
	for i := l; i > 0; i-- {
		n *= i
	}
	if l == 1 {
		return []string {s}
	}

	idx := [2]int {0, 1}
	ret := make([]string, n)
	ret_sli := make([][]string, 0, n)
	ret_sli = append(ret_sli, strings.Split(s, ""))
	for {
		add := make([][]string, 0, (l -1) * len(ret_sli))
		for {
			for _, sli := range ret_sli {
				add = append(add, swap(idx[0], idx[1], sli))
			}
			if idx[1] +1 == l {
				ret_sli = append(ret_sli, add...)
				break
			} else {
				idx[1]++
			}
		}
		if idx[0] +1 == idx[1] {
			break
		} else {
			idx[0]++
			idx[1] = idx[0] + 1
		}
	}
	for pos, sli := range ret_sli {
		ret[pos] = strings.Join(sli, "")
	}
	return ret
}

func GetCombination(s string) []string {
	dup_sli := GetPermutation(s)
	sort.Strings(dup_sli)
	unique_sli := make([]string, 0, len(dup_sli))
	unique_sli = append(unique_sli, dup_sli[0])
	before := dup_sli[0]
	for i := 1; i < len(dup_sli); i++ {
		if dup_sli[i] == before {
			continue
		}
		unique_sli = append(unique_sli, dup_sli[i])
		before = dup_sli[i]
	}
	return unique_sli
}