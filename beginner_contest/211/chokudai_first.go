package main

import (
	"fmt"
	"os"
	"strings"
)
var sum uint64
func main() {
	// Inputの内、chokudai の Ｃについて
	//   次の文字ｈがないときはゼロ
	//   ｈがあったとき、カウントを取得して、ループ開始し、最初のｈ以降の文字を取得
	//     Ｏがないときはゼロ
	// 　　Ｏがあったとき、
	//       ...
	// 　　　　　Aがないときはゼロ
	// 　　　　　Aがあったとき
	// 　　　　　　A　×　I　を総和に追加　


	// 問題点　ループ内でresultがhchokudai >> hokudaiと変化してほしいが、そうなっていない
	var s string
	sum = 0
	fmt.Scan(&s)
	c_count := strings.Count(s, "c")
	c_result := splWithFirst(s, "c")
	for h := 0; h < c_count; h++ {
		h_result := spl(c_result, "h")
		if h_result == "zero" {
			fmt.Println(0)
			os.Exit(0)
		}
		h_count := strings.Count(h_result, "h")
		for i := 0; i < h_count; i++ {
			o_result := spl(h_result, "o")
			if o_result == "zero" {
				fmt.Println(0)
				os.Exit(0)
			}
			o_count := strings.Count(o_result, "o")
			for j := 0; j < o_count; j++ {
				k_result := spl(o_result, "k")
				if k_result == "zero" {
					fmt.Println(0)
					os.Exit(0)
				}
				k_count := strings.Count(k_result, "k")
				for k := 0; k < k_count; k++ {
					u_result := spl(k_result, "u")
					if u_result == "zero" {
						fmt.Println(0)
						os.Exit(0)
					}
					u_count := strings.Count(u_result, "u")
					for l := 0; l < u_count; l++ {
						d_result := spl(u_result, "d")
						if d_result == "zero" {
							fmt.Println(0)
							os.Exit(0)
						}
						d_count := strings.Count(d_result, "d")
						for m := 0; m < d_count; m++ {
							a_result := spl(d_result, "a")
							if a_result == "zero" {
								fmt.Println(0)
								os.Exit(0)
							}
							a_count := strings.Count(a_result, "a")
							for n := 0; n < a_count; n++ {
								i_result := spl(a_result, "i")
								sum += uint64(strings.Count(i_result, "i"))
								a_result = spl(a_result, "a")
							}
							d_result = spl(d_result, "d")
						}
						u_result = spl(u_result, "u")
					}
					k_result = spl(k_result, "k")
				}
				o_result = spl(o_result, "o")
			}
			h_result = spl(h_result, "h")
		}
		c_result = spl(c_result, "c")
	}
	fmt.Println(sum % (1000000000 + 7))
}

// 文字列baseを受け取り、目的の文字(c)よりあとの文字列を返す
func spl(base string, c string) string {
	base = base[1:]
	for pos, i := range base {
		if string(i) == c {
			return base[pos:]
		}
	}
	return "zero"
}

func splWithFirst(base string, c string) string {
	for pos, i := range base {
		if string(i) == c {
			return base[pos:]
		}
	}
	return "zero"
}
