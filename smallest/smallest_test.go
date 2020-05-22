package smallest



import "testing"

var testCases = []struct {
	A []int
	expected int
}{
	{[]int{1, 3, 6, 4, 1, 2}, 5},
	{[]int{1}, 2},
	{[]int{-1}, 1},
	{[]int{-1, 0, 1}, 2},
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
