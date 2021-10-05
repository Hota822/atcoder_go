// ========================  combination string  =============================
// e.g. 123 => [123, 132, 213, 231, 312, 321]
import (
	"fmt"
	"strings"
)

func GetCombination(s string) {
	swap := func(i, j int, sli []string) []string {
		new := make([]string, len(sli))
		copy(new, sli)
		new[i], new[j] = sli[j], sli[i]
		return new
	}

	var n int = 1
	l := len(s)
	for i := l; i > 0; i-- {
		n *= i
	}
	idx := [2]int {0, 1}
	ret := make([]string, n)
	ret_sli := make([][]string, 0, n)
	ret_sli = append(ret_sli, strings.Split(s, ""))
	for {
		add := make([][]string, 0, (l -1) * len(ret_sli))
		for {
			for _, sli := range ret_sli {
				add = append(add, swap(idx[0], idx[1], sli))
			}
			if idx[1] +1 == l {
				ret_sli = append(ret_sli, add...)
				break
			} else {
				idx[1]++
			}
		}
		if idx[0] +1 == idx[1] {
			break
		} else {
			idx[0]++
			idx[1] = idx[0] + 1
		}
	}
	for pos, sli := range ret_sli {
		ret[pos] = strings.Join(sli, "")
	}
	return ret
}