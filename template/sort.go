package sort

// integer
func Ints(slice []int) []int {
	// return sorted slice
}

// struct====================================================================================
// require 3 methods, len() , swap(), less() to the struct
// and declare type, slice made by struct
// e.g. order by ascending

type card struct {
	num    int
	coords int
}

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i].coords < a[j].coords }

type SortBy []card

// usage
// func RunSort() {
// 	sli := make([]card, n)
// 	a_sli[i] = card{i: i, val: readInt()}
// 	sort.Sort(SortBy(as))
// }

// sample-----------------------
// sort []string (length is constant)
func (a SortBy) Less(i, j int) bool {
	ai, aj := a[i], a[j]
	var x int
	for x = 0; i < sli_len; i++ {
		if ai[x] == aj[x] {
			continue
		}
		break
	}
	return ai[x] < aj[x]
}

// topological sort
