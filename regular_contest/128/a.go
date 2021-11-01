package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	// "strings"
	// "runtime"
	"math"
	// "reflect"
	// "sort"
)

const (
    max_bufSize = 1_000_000_000 // default: 65536
	initial_buf = 10000
	// max_int32 = 2147483647
	max_order = 1_000_000_000
	// prime_number = 1000_000_007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	n := readInt()
	// s := read()

	sli := readSli(n)

	// dp := make([][]Gram, n +1)
	dp := make([][]float64, n +1)
	for i:=0; i <= n; i++ {
		dp[i] = make([]float64, 2)
	}

	// gold = 0, silver = 1
	dp[0][0] = math.Log10(1.0)
	// dp[0][0] = Gram {over: 0, g: 1.0}
	dp[0][1] = 0.0
	// dp[0][1] = Gram {over: 0, g: 0.0}
	ans := make([][]int, n +1)
	// ans[0] = make([]int, 2)
	for i:=1; i <= n; i++ {
		ans[i] = make([]int, 2)
		rate := float64(sli[i -1])
		before := dp[i -1]
		// gold = max (before_silver / rate,  before_gold)
		dp[i][0], ans[i][0] = Max(before[1] - math.Log10(rate), before[0])
		// silver = max(before_gold * rate before_gold)
		dp[i][1], ans[i][1] = Max(before[0] + math.Log10(rate), before[1])
	}

	ans_sli := make([]int , n)
	ans_sli[0] = 100
	idx := 0
	// fmt.Println(ans)
	for i:=n -1; i >= 0; i-- {
		// fmt.Println(ans_sli[i])
		// fmt.Println(ans[i +1][idx % 2])
		// fmt.Println(i)
		ans_sli[i] = ans[i +1][idx % 2]
		if ans_sli[i] == 1 {
			idx ++
		}
	}
	// ans := dp[n][0]
	// fmt.Println("=====================")
	return ans_sli
}

func Max(x, y float64) (float64, int){
	if x > y {
		return x, 1
	}
	return y, 0
}

// type Gram struct {
// 	over int
// 	g float64
// }

// func MaxSilver(gold Gram, rate int, silver Gram) Gram {
// 	new := Gram {over: 0, g: 0}
// 	v1 := (gold.g * rate) / max_order
// 	if v1 < 1 {
// 		new.over += 1
// 		new.g = v1 - max_order
// 	} 
// }

// func MaxGold(silver Gram, rate int, gold Gram) Gram {
// 	 new := Gram {over: 0, g: 0}

// 	 v2 := silver.over * max_order / float64(rate)
// 	 if v2 < max_order {
// 		v3 := silver.g + v2
// 		if v3 > max_order {
// 			new.over = 1
// 			new.g = v3 - max_order
// 		} else {
// 			new.g = v3
// 		}
// 	 } else {
// 		 new.over := int(silver.over * float64(rate))
// 		 v3 := v2 % max_order
// 		 new.g = v3
// 	 }

// 	return Max(new, gold)
// }

// func Max(x, y Gram) Gram {
// 	is x.over == y.over {
// 		if x.g > y.g {
// 			return x
// 		} else {
// 			return y
// 		}
// 	}

// 	if x. over > y.over {
// 		return x
// 	}
// 	return y
// }


// ========================read
// func read() string {
// 	sc.Scan()
//     return sc.Text()
// }

// func readSli(n int) []string {
func readSli(n int) []int {
	// sli := make([]string, n)
	sli := make([]int, n)
	for i:=0; i < n; i++ {
		// sli[i] = read()
		sli[i] = readInt()
	}
    return sli
}

func readInt() int {
    sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
    return ret
}


//=======================main========================
func main() {
	buf := make([]byte, initial_buf)
	sc.Buffer(buf, max_bufSize)
	sc.Split(bufio.ScanWords)
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
		l := len(sli)
		for i:=0; i < l; i ++ {
			fmt.Print(sli[i])
			fmt.Print(" ")
		}
		fmt.Println("")
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

// func d(arg ...interface{}) {
// 	_, _, l, _ := runtime.Caller(1)
// 	s := strconv.Itoa(l)
// 	fmt.Println("dumped at line: " + s + ", value: ")
// 	for _, v := range arg {
// 		fmt.Println(v)
// 	}
// }
