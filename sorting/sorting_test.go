package smallest



import "testing"

var testCases = []struct {
	A int
	B int
	K int
	expected int
}{
	{5, 12, 2, 4},
	{6, 12, 2, 4},
	{6, 13, 3, 3},
	{6, 11, 2, 3},
	{7, 11, 2, 2},
	{17, 21, 2, 2},
	{6, 11, 7, 1},
	{6, 21, 7, 3},
	{6, 49, 7, 7},
	{2, 2, 2, 1},
	{0, 0, 1, 1},
	{0, 1, 1, 2},
	{0, 1, 2, 1},
	{0, 1, 11, 1},
	{10, 10, 7, 0},
	{10, 10, 22, 0},
	{10, 10, 5, 1},
}

func TestSolution(t *testing.T) {
	// Given
	for _, testCase := range testCases {
		// WhenR
		result := Solution(testCase.A, testCase.B, testCase.K)

		if testCase.expected != result {
			t.Errorf("Test failed input: A = %d expected: %d, received: %d", testCase, testCase.expected, result)
		}
	}
}
