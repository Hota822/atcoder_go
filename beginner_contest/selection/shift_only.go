package main
import  (
    "fmt"
    "bufio"
    "os"
    "strconv"
    // "reflect"
)
var sc = bufio.NewScanner(os.Stdin)
func main() {
    sc.Split(bufio.ScanWords)
    var times int
    times = 0
    total := nextInt()
    strings := make([]string, total)
    lengths := make([]int, total)
    // 対象を全て２進数に変換して保持し、その長さも保持する
    for  i := 0; i < total; i++ {
        strings[i] = fmt.Sprintf("%b", nextInt())
        lengths[i] = len(strings[i])
    }
    // 末尾から１文字ずつ取得し、その総和を計算。１なら終了、０なら続行
    for j := 0; j < 1000000000; j++ {
        var result bool
        result = dividable(j, lengths, strings)
        if (!result) {
            fmt.Println(times)
            return
        }
        times += 1
    }
    // fmt.Println("aaaa")
}
func nextInt() int {
    sc.Scan()
    ret, _ := strconv.Atoi(sc.Text())
    return ret
}
func dividable(cool int, lengths []int, strings []string) bool {
    // defer func() {
    //     // return false
    // }()
    var sum int
    for pos, k := range strings {
        ret, _ := strconv.Atoi(string(k[lengths[pos] - cool - 1]))
        sum += ret
    }
    fmt.Println(sum)
    return sum == 0
}