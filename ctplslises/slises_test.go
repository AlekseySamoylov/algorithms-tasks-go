package ctplslises

import "testing"

func TestSolution(t *testing.T) {
	// Given
	M := 78
	A := []int{3,4,5,6,2,6,3,78,5,3,4,0,9,6,6,6,4,5,4,32,21,1}
	//A := []int{0}
	expectedResult := Solution(M, A)

	// When
	result := Solution_2(M, A)

	if expectedResult != result {
		t.Errorf("Test failed expected: %d, received: %d", expectedResult, result)
	}
}

/*
1 -> 1
2 -> 3
3 -> 6
4 -> 10
5 -> 15
6 -> 21
 */
