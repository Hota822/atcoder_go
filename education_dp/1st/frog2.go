package main
// 1:12:05
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	// "runtime"
	"math"
	// "reflect"
	// "sort"
	// "strings"
)

const (
    // max_bufSize = 100000 // default: 64000
	// prime_number = 1000000007
)

var sc = bufio.NewScanner(os.Stdin)

func run() interface{} {
	// buf := make([]byte, maxBufSize)
	// sc.Buffer(buf, max_bufSize)
	sc.Split(bufio.ScanWords)

	n, k := readInt(), readInt()
	sli := make([]int, n)
	cost := make([]int, n)
	sli[0] = readInt()
	cost[0] = 0
	for i:=1; i < n; i++ {
		h := readInt()
		sli[i] = h
		// min := make([]int, 0, k)
		min := 2147483647
		for j:=1; j <= k; j++{
			idx := i - j
			
			if idx < 0 {
				break
			}
			ci := DiffAbs(h, sli[idx])
			sum_cost := ci + cost[idx]
			if sum_cost < min {
				min = sum_cost
			}
			// d(min)
			// min_sli = append(min_sli, DiffAbs(h, sli[idx]))
		}
		cost[i] = min
		// d(cost)
	}
	// sli := make([]int, n)
	// sli := make([]string, n)

	return cost[n -1]
}

func DiffAbs(x, y int) int {
	return int(math.Abs(float64(x - y)))
}

// func MinSli(sli []int) int {
// 	sort.Ints(sli)
// 	return sli[0]
// }


// ========================read
func readInt() int {
    sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
    return ret
}


//=======================main========================
func main() {
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

// func d(arg ...interface{}) {
// 	_, _, l, _ := runtime.Caller(1)
// 	s := strconv.Itoa(l)
// 	fmt.Println("dumped at line: " + s + ", value: ")
// 	for _, v := range arg {
// 		fmt.Println(v)
// 	}
// }