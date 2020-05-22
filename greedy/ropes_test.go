package greedy



import "testing"

var testCases = []struct {
	K int
	A []int
	expected int
}{
	{4, []int{1,2,3,4,1,1,3}, 3},
	{6, []int{1,2,3,2,1,1,3}, 2},
	{2, []int{1,2,3,2,1,1,3}, 5},
	{4, []int{4,5}, 2},
	{3, []int{1,3}, 1},
	{4, []int{6}, 1},
}

func TestSolution(t *testing.T) {
	// Given
	for _, testCase := range testCases {
		// WhenR
		result := Solution(testCase.K, testCase.A)

		if testCase.expected != result {
			t.Errorf("Test failed input: K = %d A = %d expected: %d, received: %d", testCase.K, testCase.A, testCase.expected, result)
		}
	}
}
