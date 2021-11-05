package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"runtime"
	"strings"
	// "math"
	// "reflect"
	// "sort"
)

const (
    max_bufSize = 1_000_000_000 // default: 65536
	initial_buf = 10000
	// max_int32 = 2147483647
	// prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	// n := readInt()
	s, t := read(), read()

	// sli := make([][]int, n)
	// for i:=0; i < n; i++ {
	// 	sli[i] = readSli(2)
	// }
	s_sli := strings.Split(s, "")
	t_sli := strings.Split(t, "")
	dp := make([][]string, len(t) +1)
	m := GetAlphabetMap()
    dp[0] = make([]string, len(s) + 1)
    // for i:=0; i < len(s); i++ {
    //     m[s_sli[i]]++
    // }
    // ans := make([]string, Max(len(s), len(t)))
    for i:=1; i <= len(t_sli); i++ {
        dp[i] = make([]string, len(t) + 1)
        t_c := t_sli[i -1]
        m[t_c]++
        count := GetAlphabetMap()
        cache := ""
        for j:=1; j <= len(s_sli); j++ {
            s_c := s_sli[j -1]
            if s_c == t_c {
                // p(count[t_c], m[t_c])
                if count[t_c] < m[t_c] {
                    cache = dp[i][j -1] + t_c
                    // p(s_c, cache, i, j)
                }

                if len(cache) < len(dp[i -1][j]) {
                    count[t_c]--
                    cache = dp[i -1][j] + t_c
                    // p("if")
                }

                dp[i][j] = cache
                count[t_c]++
            } else {
                dp[i][j] = Longer(dp[i][j -1], dp[i -1][j])
            }
        }
    }

	// ans := dp[len(s)][len(t)]
    ans := dp[len(t)][len(s)]
    // dp[1][1] = "23"
    // p(dp)
	return ans
}

func Longer(s, t string) string {
    ls := len(s)
    lt := len(t)
    if ls >= lt {
        return s
    }

    return t
}

func GetAlphabetMap() map[string]int {
	m := make(map[string]int)
	alphabets := []string {"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	for i:=0; i < 26; i++ {
		m[alphabets[i]] = 0
	}
	return m
}


// ========================read
func read() string {
	sc.Scan()
    return sc.Text()
}

// func readSli(n int) []string {
// func readSli(n int) []int {
// 	// sli := make([]string, n)
// 	sli := make([]int, n)
// 	for i:=0; i < n; i++ {
// 		// sli[i] = read()
// 		sli[i] = readInt()
// 	}
//     return sli
// }

// func readInt() int {
//     sc.Scan()
//     ret, _ := strconv.Atoi(sc.Text())
//     return ret
// }


//=======================main========================
func main() {
	buf := make([]byte, initial_buf)
	sc.Buffer(buf, max_bufSize)
	sc.Split(bufio.ScanWords)
	result := run()
	Print(result)
}

func Print(ans interface{}) {
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
	if sli, ok := ans.([]string); ok {
		for _, v := range sli {
			fmt.Println(v)
		}
		return
	}
	fmt.Println(ans)
}

func PrintOne(ans interface{}) {
	if sli, ok := ans.([]int); ok {
		for _, v := range sli {
			fmt.Print(v)
		}
	}
	if sli, ok := ans.([]string); ok {
		for _, v := range sli {
			fmt.Print(v)
		}
	}
	fmt.Println("")
}

// =============================math
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return - x
}

func p(arg ...interface{}) {
	_, _, l, _ := runtime.Caller(1)
	s := strconv.Itoa(l)
	fmt.Println("dumped at line: " + s + ", value: ")
	for _, v := range arg {
		fmt.Println(v)
	}
}
