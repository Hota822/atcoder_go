
package main
 
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)
 
const inf = 1_000_000_000 + 7
 
// const createSize = 4080000
const createSize = 10000005 //ここのメモリ量で速度が定数倍変わるので気をつけるように
 
func main() {
	scan_init()
	n := scanInt()
	abc := make([][]int, 3)
	for i:=0; i<3; i++ {
		abc[i] = make([]int, n)
	}
	for j:=0; j<n; j++ {
		for i:=0; i<3; i++ {
			abc[i][j] = scanInt()
		}
	}
	dp := make([][]int, n+1)
	for i:=0; i<=n; i++ {
		dp[i] = make([]int, 3)
	}
	for i:=0; i<3; i++ {
		dp[0][i] = abc[i][0]
	}
	for i:=1; i<n; i++ {
		for j:=0; j<3; j++ {
			me := abc[j][i]
			output := me
			for k:=0; k<3; k++ {
				if j == k {
					continue
				}
				output = max(me+dp[i-1][k], output)
			}
			dp[i][j] = output
		}
	}
	output := 0
	for i:=0; i<3; i++ {
		output = max(output, dp[n-1][i])
	}
	fmt.Println(output)
}
 
var sc = bufio.NewScanner(os.Stdin)
func abs(a int) int{
	if a<0 {
		return -a
	} else {
		return a
	}
}
func max (a,b int) int{
	if a > b {
		return a
	}  else {
		return b
	}
}
func scan_init() {
	sc.Buffer([]byte{}, math.MaxInt64)
	sc.Split(bufio.ScanWords)
}
func scanInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}
func scanInts(n int) []int {
	take := make([]int, n)
	for i := 0; i < n; i++ {
		take[i] = scanInt()
	}
	return take
}
func scanIntTwoDimension(h int, w int) [][]int {
	take := make([][]int, h)
	for i := 0; i < h; i++ {
		take[i] = scanInts(w)
	}
	return take
}
func scan() string {
	sc.Scan()
	return sc.Text()
}
 
var rdr = bufio.NewReaderSize(os.Stdin, 200000)
 
func readLine() string {
	buf := make([]byte, 0, 200000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
func readLineToNumber() int {
	S := readLine()
	number, _ := strconv.Atoi(S)
	return number
}
func readLineToSlice() []string {
	S := readLine()
	list := strings.Split(S, "")
	return list
}
func readLineToNumberSlice(memo map[string]int) []int {
	// err時は-1を代入
	S := readLine()
	intList := make([]int, len(S))
	for i := 0; i < len(S); i++ {
		if val, ok := memo[string(S[i])]; ok {
			intList[i] = val
		} else {
			intList[i] = -1
		}
	}
	return intList
}
 
type Deque struct {
	cl   int
	cr   int
	data []interface{}
}
 
func (d *Deque) push(x interface{}) {
	d.cl--
	d.data[d.cl] = x
}
 
func (d *Deque) unshift(x interface{}) {
	d.data[d.cr] = x
	d.cr++
}
 
func (d *Deque) pop() interface{} {
	v := d.data[d.cl]
	d.cl++
	return v
}
 
func (d *Deque) shift() interface{} {
	v := d.data[d.cr-1]
	d.cr--
	return v
}
 
func (d *Deque) size() int {
	return d.cr - d.cl
}
 
func (d *Deque) empty() bool {
	return d.size() == 0
}
 
func (d *Deque) get(x int) interface{} {
	return d.data[d.cl+x-1]
}
 
func createDeque(size int) Deque {
	// 取り出すときは、アサーション忘れずに
	// i.(type)　こういうやつ
	return Deque{cl: size, cr: size, data: make([]interface{}, size*2, size*2)}
}