package print_slice

func CreateSlice(n int) (result [][]int) {
	result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	return result
}
