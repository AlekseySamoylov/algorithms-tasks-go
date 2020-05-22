package fishstream



import "testing"

var testCases = []struct {
	A []int
	B []int
	expected int
}{
	{[]int{4,3,2,1,5}, []int{0,1,0,0,0}, 2},
	{[]int{4,3,2,1,5}, []int{0,1,1,1,0}, 2},
	{[]int{4,3,2,1,0}, []int{0,1,1,1,0}, 4},
	{[]int{7,4,2,6,1,3,5}, []int{1,1,1,1,0,0,0}, 4},
	{[]int{7,4,2,6,2,1,3,5}, []int{1,1,1,1,1,0,0,0}, 4},
	{[]int{4,3,2,1,5,4,3,2,1,5}, []int{0,1,1,1,0,0,1,1,1,0}, 4},
	{[]int{4,3}, []int{0,1}, 2},
	{[]int{4,3}, []int{1,0}, 1},
	{[]int{4}, []int{1}, 1},
	{[]int{4}, []int{0}, 1},
	{[]int{4,3,2,0,5}, []int{0,1,0,0,0}, 2},

}

func TestSolution(t *testing.T) {
	// Given
	for _, testCase := range testCases {
		// WhenR
		result := Solution(testCase.A, testCase.B)

		if testCase.expected != result {
			t.Errorf("Test failed input: A = %d B = %d expected: %d, received: %d", testCase.A, testCase.B, testCase.expected, result)
		}
	}
}
