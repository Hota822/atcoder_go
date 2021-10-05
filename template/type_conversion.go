// string ----> int(n)------------------------------
import (
	"strconv"
)
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

// function
func ConvertInt64(s string) int64 {
	ret, _ = strconv.ParseInt(s, 10, 64)
	return ret
}

// int(n) ----> string------------------------------
import (
	"strconv"
)
// int64
func Convert64ToString(i int64) string {
	ret := strconv.FormatInt(n, 10)
	return ret
}
// int
func ConvertToString(i int) string {
    ret := strconv.Itoa(i)
    return ret
}