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
	var n, k int
	sc.Split(bufio.ScanWords)
	n, k = nextInt(), nextInt()
	if k == 1 {
		fmt.Println(1)
		return
	}
	c_sli := make([]int, n)
	for i := 0; i < k; k++ {
		c_sli[i] = nextInt()
	}

	max := 1
	types := make(map[int] cnt, k)
	for i := 0; i < k; i++ {
		next := c_sli[i]
		_, yn1 := types[next]
		if yn1 {
			types[next] = cnt {count: types[next].count + 1}
			// types[next].count += 1
		} else {
			types[next] = cnt {count: 1}
		}
	}
	for i := 0; i < n - k + 1; i++ {
		next := c_sli[i + k-1]
		_, yn1 := types[next]
		if yn1 {
		types[next] = cnt {count: types[next].count + 1}
		} else {
			types[next] = cnt {count: 1}
		}

		length := len(types)
		if length > max {
			max = length
		}

		first := c_sli[i]
		count := types[first].count
		if count == 1 {
			delete(types, first)
		} else {
			count -= 1
		}
	}
	fmt.Println(max)
}

type cnt struct {
	count int
}

func nextInt() int {
    sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
    return ret
}