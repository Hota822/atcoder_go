func initSliceAs(len, val int) []int {
	sli := make([]int, len)
	sli[0] = val
	for i := 1; i < len; i *= 2 {
		copy(sli[i:], sli[:i])
	}
	return sli
}
