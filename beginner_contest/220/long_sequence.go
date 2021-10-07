package main
​
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)
​
var sc = bufio.NewScanner(os.Stdin)
func main() {
    sc.Split(bufio.ScanWords)
    n := next()
    sum := make([]int64, n)
​
    sum[0] = next64()
​
    for i := 1; i < n; i++ {
        sum[i] = sum[i -1] + next64()
    }
    x := next64()
    loop, remain := ForA(sum[n -1], x)
    ans := loop * int64(n)
​
    for _, i := range sum {
        ans++
        if remain < i {
            break
        }
    }
    fmt.Println(ans)
}
​
func next() int {
    sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
    return ret
}
​
​
func next64() int64 {
    sc.Scan()
    ret, _ := strconv.ParseInt(sc.Text(), 10, 64)
    return ret
}
​
func ForA(sum, x int64) (int64, int64) {
    count := x / sum
    remain := x % sum
    return count, remain
}