package main

import (
	"fmt"
	"os"
	"strings"
)
var sum uint64
func main() {
	var s string
	sum = 0
	fmt.Scan(&s)
	forChr(s, 0)
	fmt.Println(sum % (1000000000 + 7))
}

// 文字列baseを受け取り、目的の文字(c)よりあとの文字列を返す
func spl(base string, c string) string {
	for pos, i := range base {
		if string(i) == c {
			return base[pos:]
		}
	}
	return "zero"
}
// Inputの内、chokudai の Ｃについて
//   次の文字ｈがないときはゼロ
//   ｈがあったとき、
//     Ｏがないときはゼロ
// 　　Ｏがあったとき、
//       ...
// 　　　　　Ａがないときはゼロ
// 　　　　　次の文字がＩの時はＩの個数をカウントして総和に追加（関数でまとめるために分岐追加）


func forChr(str string, current int) {
	fmt.Println(current, str)
	s := "chokudai"
	c_chr := string(s[current])
	n_chr := string(s[current + 1])
	count := strings.Count(str, c_chr)
	// fmt.Println(count)
	for n := 0; n < count; n++ {
		result := spl(str, n_chr)
		if result == "zero" {
			fmt.Println(0)
			os.Exit(1)
		} else if n_chr == "i" {
			count_i := uint64(strings.Count(result, "i"))
			sum += count_i
			continue
		}
		forChr(result, current + 1)
	}
	// CHchokudai
	// Hchokudai
	// okudai

	// CHchokudai
	// hokudai

	// chCHokudai
	// ChcHokudai
}
