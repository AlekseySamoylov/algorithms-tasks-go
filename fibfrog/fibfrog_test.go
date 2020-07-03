package fibfrog

import "testing"

var testCases = []struct {
	A        []int
	expected int
}{
	{[]int{0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1}, 2},
	{[]int{0, 0, 0, 1, 1, 0, 1, 0, 0, 0, 0}, 3},
	{[]int{0}, 1},
	{[]int{0, 0}, 1},
	{[]int{0, 0, 0}, -1},
	{[]int{0, 0, 1}, 2},
	{[]int{0, 0, 1, 0}, 1},
	{[]int{0, 0, 1, 0, 0, 0}, -1},
	{[]int{0, 0, 1, 0, 0, 1, 0, 1}, 2},
	{[]int{0, 0, 1, 0, 0, 1, 0}, 1},
	{[]int{0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 0, 1, 0}, 4},
	{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 2},
}

func TestSolution(t *testing.T) {
	// Given
	for _, testCase := range testCases {
		// When
		result := Solution(testCase.A)

		if testCase.expected != result {
			t.Errorf("Test failed input: A = %d expected: %d, received: %d", testCase.A, testCase.expected, result)
		}
	}
}