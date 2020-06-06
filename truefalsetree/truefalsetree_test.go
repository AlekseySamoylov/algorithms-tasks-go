package truefalsetree

import "testing"

var testCases = []struct {
	A        []bool
	expected int
}{
	{[]bool{true, false, false, true, false}, 11},
	{[]bool{true, false, false, true}, 7},
	{[]bool{true}, 1},
	{[]bool{true, true, true, true}, 10},
	{[]bool{true, false, false, false}, 4},
	{[]bool{false, false, false, false, false, false, true}, 7},
	{[]bool{false}, 0},
	{[]bool{false, false, false, false, false}, 0},
}

func TestSolution(t *testing.T) {
	// Given
	for _, testCase := range testCases {
		// When
		result := Solution(testCase.A)

		if testCase.expected != result {
			t.Errorf("Test failed input: A = %t expected: %d, received: %d", testCase.A, testCase.expected, result)
		}
	}
}
