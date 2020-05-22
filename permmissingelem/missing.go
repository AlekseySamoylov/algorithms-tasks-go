package permmissingelement

func Solution(A []int) int {
	lenA := len(A)
	if lenA == 0 {
		return 1
	}
	checker := make([]int, lenA + 2)
	checker[0] = 1
	for idx := 0; idx < lenA; idx++ {
		checker[A[idx]] = 1
	}
	for idx := 0; idx < len(checker); idx++ {
		result := checker[idx]
		if result == 0 {
			return idx
		}
	}
	return 0
}

