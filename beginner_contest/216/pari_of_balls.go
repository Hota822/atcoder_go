package main

import (
	"fmt"
	"bufio"
	"os"
	// "strings"
	"strconv"
	// "reflect"
	// "math"
	// "sort"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	// [tube] = idx
	next_sli := make([]int, m)
	// [tube][...color]
	all := make([][]int, m)
	mapping := make([]Pair,  n + 1)
	for i := 0; i < m; i++ {
		k := nextInt()
		tube := make([]int, k)
		for j := 0; j < k; j++ {
			color := nextInt()
			if len(mapping[color]) == 0 {
				mapping[color] = Pair {}
			}
			mapping[color] = append(mapping[color], Ball {tube: i, idx: j})
			tube[j] = color
		}
		next_sli[i] = k - 1
		all[i] = tube
	}
	target := 0
	for {
		before := n
		// i: number of tubes
		
		for i := 0 + offset; i < m; i ++ {
			idx := next_sli[i]
			current_color := all[i][idx]
			pair := mapping[current_color].GetPair(current_color, i, idx)
			// ペアで取り除けるか確認
			if next_sli[pair.tube] == pair.idx {
				// インデックスを次に進める処理
				if idx == 0 {
					all[i][0] = 0
				} else {
					next_sli[i] -= 1
				}
				if next_sli[pair.tube] == 0 {
					all[pair.tube][0] = 0
				} else {
					next_sli[pair.tube] -= 1
				}

				n -= 2
				// ペアで取り除けたら初めに戻る
				break
			}
		}
		if n == 0 {
			fmt.Println("Yes")
			return
		}
		if before == n {
			fmt.Println("No")
			return
		}
	}
}

type Pair []Ball

type Ball struct {
	idx int
	tube int
}

func (p Pair) GetPair(color int, tube, idx int) Ball{
	if p[0].tube == tube {
		if p[0].idx == idx {
			return p[1]
		}
	}
	return p[0]
}

// func next() string {
// 	sc.Scan()
//     ret := sc.Text()
//     return ret
// }

func nextInt() int {
    sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
    return ret
}

// func nextInt64() int64 {
// 	var ret int64
//     sc.Scan()
//     ret, _ = strconv.ParseInt(sc.Text(), 10, 64)
//     return ret
// }