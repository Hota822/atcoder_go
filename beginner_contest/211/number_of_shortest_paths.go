package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	"strconv"
	// "reflect"
	// "math"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	// var x int
	sc.Split(bufio.ScanWords)
	// fmt.Scan(&x)
	target_city := nextInt()
	loads_count := nextInt()
	atob := make([][]int, loads_count)
	for i := 0; i < loads; i++ {
		a := []int {nextInt(), nextInt(), 0}
		atob[i] = a
	}

	got_cities := make(map[int]struct{})
	got_cities[target_city] = struct{}
	next_cities := make(map[int]struct{})

	// ゴールにたどり着くものがある時、それを次の目的地とする
	for j := 0; j < loads_count; j++ {
		if atob[j][1] == target_city {
			next_cities[j] = struct{}
		}
	}
	for {
		// 差分をとる
		next_cities = diff_cities(next_cities)
		// たどり着くものから、ペアを取得
		post_next_cities := make(map[int]struct{})
		for _, c := range next_cities
			for k := 0; k < loads_count; k++ {
				if atob[k][0] == c {
					post_next_cities[k] = struct{}
				} else if atob[k][1] {
					post_next_cities[k] = struct{}
				}
			}
		}
		next_cities = post_next_cities
	}
	// fmt.Println(len(array))
	fmt.Println("end", target, loads, array, atob)
}

// 	// got_cities := make([][]int, loads_count)
// 	got_cities := make(map[int]struct{})
// 	// got_cities[1] = struct{}
// 	for {
// 		// 次の場所を初期化
// 		next_cities := make(map[int]struct{})
// 		// atobの１つずつを確認
// 		for j := 0; j < loads_count; j++ {
// 			// 目的地についた場合、ループ回数を出力し
// 			if atob[j][1] == target_city {

// 			}
// 			// すでにいったことがあるかを確認
// 			_, exists_a := got_cities[atob[j][0]];
// 			_, exists_b := got_cities[atob[j][1]];

// 			// どちらか片方いったことない場合、 次の目的地に追加する
// 			if  exists_a || exists_b {
// 				got_cities[atob[j][1]] = struct{}
// 				next_cities[atob[j][1]] = struct{}
// 			}
// 			// got_cities[0] := make([]int, loads_count)
// 			if atob[j][0] == 1 {
// 				// got_cities[atob[j][1]] = 1
// 			}
// 		}

// 		if len(next_cities) == 0 {
// 			fmt.Println(0)
// 		}
// 	}

	// for _, i := range array {
	// 	fmt.Println(i)
	// }
	// for i := 0; i < loads_count; i++ {
	// 	set := make(map[int]struct{})
	// 	if atob[i][1] == target {
	// 		atob[2] += 1 // 使用回数追加
	// 		set[i] = i // 次の数の取得
	// 	}
	// 	for j, v := range set {
	// 		if atob[i][1] == target {
	// 			atob[2] += 1 // 使用回数追加
	// 			set[i] = i // 次の数の取得
	// 		}
	// 	}
	// 	for j := 0; j < len(next_targets); j ++ {
	// 		if atob[i][1] ==  {
	// 			array[i] += 1
	// 			next_targets.append(i)
	// 		}
	// 	}
	// }
	// next_targets:= make([]int, 0, loads)
	// if atob[i][1] == target {
	// 	array[i] += 1
	// 	next_targets.append(i)
	// }
	// for j := 0; j < len(next_targets); j ++ {
	// 	if atob[i][1] ==  {
	// 		array[i] += 1
	// 		next_targets.append(i)
	// 	}
	// }
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