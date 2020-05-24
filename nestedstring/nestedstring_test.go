package nestedstring

import "testing"

var testCases = []struct {
	A        string
	expected bool
}{
	{"()()()", true},
	{"()", true},
	{"(()()())", true},
	{"(()(()))((()))(()())", true},
	{"(()()))((()))(()())", false},
	{"(()()))", false},
	{"(", false},
	{"(()(()))((()))(()()))", false},
	{"(()(()))((()))(()()", false},
	{"))((", false},
}

func TestSolution(t *testing.T) {
	// Given
	for _, testCase := range testCases {
		// WhenR
		result := Solution(testCase.A)

		if testCase.expected != result {
			t.Errorf("Test failed input: A = %s expected: %t, received: %t", testCase.A, testCase.expected, result)
		}
	}
}
