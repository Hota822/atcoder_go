package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	var s, t string
	fmt.Scan(&s)

    for _, c := range s {
		if strings.Count(s, string([]rune{c})) == 4 {
			fmt.Println("Weak")
			return
		}
		t = string(c) + next(string(c))
		// t = string([]rune{c}) + next(string([]rune{c}))
		break
    }

	if s == t {
		fmt.Println("Weak")
		return
	}
	fmt.Println("Strong")
}

func next(n string) string {
	var i, i1, i2, i3 int
	var j1, j2, j3 string
	i, _ = strconv.Atoi(n)
	i1 = i + 1
	i2 = i + 2
	i3 = i + 3

	j1 = strconv.Itoa(i1)
	j2 = strconv.Itoa(i2)
	j3 = strconv.Itoa(i3)
	len1 := len(j1) - 1
	len2 := len(j2) - 1
	len3 := len(j3) - 1
	return j1[len1:] + j2[len2:] + j3[len3:]
}
