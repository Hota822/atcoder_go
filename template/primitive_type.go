// -------------------map--------------------------------
// make(map[{key_type}]{value_type}, {capacity})
types := make(map[int]string, n)
// use struct {} as value is the lowest cost
types := make(map[int]struct{}, n)

// -----------------slice--------------------------------
// make([]{value_type}, {size}, {capacity} )
slice := make([]int, n)
// append
sli := make([]int, 0, n)
sli = append(sli, 1, 2, 3)
// => [1, 2, 3]
sli_2 := []int {4, 5}
sli_3 := append(sli, sli_2...)
// =>[1, 2, 3, 4, 5]
sli := make([]int, n)
a := append(sli, 1, 2, 3)
// => NG, over capacity

// copy
sli_4 := copy(sli)