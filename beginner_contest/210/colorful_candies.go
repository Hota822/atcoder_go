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



// package main

// import (
//     "fmt"
//     "bufio"
//     "os"
//     "strconv"
// )


// var sc = bufio.NewScanner(os.Stdin)     
// func main() {
//     sc.Split(bufio.ScanWords)
//     n, k := readInt(), readInt()        
//     sli := make([]int, n)

//     for i:=0; i < n; i++ {
//         sli[i] = readInt()
//     }

//     candies := make(map[int]int, n)     
//     for i:=0; i < k; i++ {
//         candy := sli[i]
//         if _, ok := candies[candy]; ok {
//             candies[candy]++
//         } else {
//             candies[candy] = 1
//         }
//     }

//     colors := len(candies)
//     for i:=0; i < n - k; i++ {
//         del_candy := sli[i]
//         get_candy := sli[i + k]

//         // delete
//         if candies[del_candy] == 1 {
//             delete(candies, del_candy)
//         } else {
//             candies[del_candy]--
//         }
//         // get
//         if _, ok := candies[get_candy]; ok {
//             candies[get_candy]++
//         } else {
//             candies[get_candy] = 1
//         }
//         colors = max(colors, len(candies))
//     }
//     fmt.Println(colors)
// }

// func readInt() int {
//     sc.Scan()
//     ret, _ := strconv.Atoi(sc.Text())
//     return ret
// }

// func max(x, y int) int {
//     if x > y {
//         return x
//     }
//     return y
// }

// // map key exists
// // mpa[key] => return value 2 is bool

// // map count
// // len(map)

// // delete key of map
// // delete(map, key)