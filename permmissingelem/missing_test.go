package permmissingelement



import "testing"

var testCases = []struct {
	A []int
	expected int
}{
	{[]int{2,3,1,5}, 4},
	{[]int{2,3,4}, 1},
	{[]int{}, 1},
	{[]int{2}, 1},
	{[]int{1,3}, 2},
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
