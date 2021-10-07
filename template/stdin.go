import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

const (
    // maxBufSize = 100000
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	// buf := make([]byte, maxBufSize)
	// sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords)
}

// scan
func ReadInt64() int64 {
	var ret int64
    sc.Scan()
    ret, _ = strconv.ParseInt(sc.Text(), 10, 64)
    return ret
}

func ReadInt() int {
	var ret int
    sc.Scan()
    ret, _ = strconv.Atoi(sc.Text())
    return ret
}


func next() string {
	sc.Scan()
    ret := sc.Text()
    return ret
}