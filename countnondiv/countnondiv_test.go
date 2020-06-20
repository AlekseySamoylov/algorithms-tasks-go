package countnondiv

import "testing"

var testCases = []struct {
	A        []int
	expected []int
}{
	{[]int{3, 1, 2, 3, 6}, []int{2, 4, 3, 2, 0}},
	{[]int{3, 1, 2, 3, 6, 3, 1, 2, 3, 6}, []int{4, 8, 6, 4, 0, 4, 8, 6, 4, 0}},
	{[]int{3}, []int{0}},
	{[]int{1, 3}, []int{1, 0}},
	{[]int{3, 3}, []int{0, 0}},
}

// N 1 .. 50,000
// a 1 .. 100,000
func TestSolution(t *testing.T) {
	// Given
	for _, testCase := range testCases {
		// WhenR
		result := Solution(testCase.A)

		if !IntArrayEquals(testCase.expected, result) {
			t.Errorf("Test failed input: A = %d expected: %d, received: %d", testCase.A, testCase.expected, result)
		}
	}
}

func IntArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
