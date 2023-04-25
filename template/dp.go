package dp

func initSliceAs(len, val int) []int {
	sli := make([]int, len)
	sli[0] = val
	for i := 1; i < len; i *= 2 {
		copy(sli[i:], sli[:i])
	}
	return sli
}

// normal tree =================================================
var tree map[int]*Node

type Node struct {
	links []int
	self  int
}

// get children
func (node *Node) getChildren(parent int) []int {
	ret := make([]int, 0)
	for _, v := range node.links {
		if v != parent {
			ret = append(ret, v)
		}
	}

	return ret
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

func (tree SegmentTree) getMax(x, y int) int {
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

func newSegmentTree(n int) SegmentTree {
	n2 := 1
	// 2^nの要素数となるように計算
	for n2 <= n {
		n2 *= 2
	}
	st := SegmentTree{element_num: n}
	st.data = make([]int, n+1)
	// p(len(st.data), n)
	return st
}

// func (tree SegmentTree) getMax(idx int) int {
// 	max := 0
// 	for i := idx; i > 0; {
// 		max = Max(max, tree.data[i])
// 		i -= i & -i
// 	}
// 	return max
// }

func (tree SegmentTree) update(idx, val int) {
	for i := idx; i <= tree.element_num; {
		tree.data[i] = Max(val, tree.data[i])
		i += i & -i
	}
}

// dependent to another packages
func Max(x, y int) int
func readInt() int
