package smallest

import "sort"

func Solution(A []int) int {
	sort.Ints(A)
	result := 1
	for idx := 0; idx < len(A); idx++ {
		if A[idx] > 0 && A[idx] == result {
			result++
		}
	}
	return result
}
