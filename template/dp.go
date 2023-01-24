func initSliceAs(len, val int) []int {
	sli := make([]int, len)
	sli[0] = val
	for i := 1; i < len; i *= 2 {
		copy(sli[i:], sli[:i])
	}
	return sli
}

var tree map[int]*Node
type Node struct {
	links []int
}

// get children
// for _, v:= range tree.links
//   if v == parent || v == current {
//     continue
//   }
//
//   any processes
// }
}

func makeTree(n int) {
	tree = map[int]*Node{}
	// tree_count := make([]int, n+1)

	createNode := func(x, y int) {
		if _, ok := tree[x]; ok {
			tree[x].links = append(tree[x].links, y)
		} else {
			tree[x] = &Node{links: []int{y}}
		}
	}
	for i := 0; i < n-1; i++ {
		x, y := readInt(), readInt()
		createNode(x, y)
		createNode(y, x)
	}
}

// segment tree =================================================
type SegmentTree struct {
	data        []int
	element_num int
}

func getMax(x, y int, tree SegmentTree) int {
	var get func(x, y, l, r, k int) int
	get = func(x, y, l, r, k int) int {
		if x <= l && r <= y { // 範囲内
			return tree.data[k]
		} else { // 一部被っているとき。範囲外は考えない
			max_l := get(x, y, l, (l+r)/2, k*2+1)
			max_r := get(x, y, (l+r)/2, r, k*2+2)
			return Max(max_l, max_r)
		}
	}
	return get(x, y, 0, tree.element_num, 0)
}
