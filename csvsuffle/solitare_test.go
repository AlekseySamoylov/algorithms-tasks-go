package scvshuffle



import "testing"

var testCases = []struct {
	A []int
	expected int
}{
	{[]int{1,-2,0,9,-1,-2}, 8},
	{[]int{1,-2,0,9,-1}, 9},
	{[]int{1,-2}, -1},
//	{[]int{1,2,3,2,1,1,3}, 2},
//	{[]int{1,2,3,2,1,1,3}, 5},
//	{[]int{4,5}, 2},
//	{[]int{1,3}, 1},
//	{[]int{6}, 1},
}

func TestSolution(t *testing.T) {
	// Given
	for _, testCase := range testCases {
		// WhenR
		result := Solution(testCase.A)

		if testCase.expected != result {
			t.Errorf("Test failed input: A = %d expected: %d, received: %d", testCase.A, testCase.expected, result)
		}
	}
}
