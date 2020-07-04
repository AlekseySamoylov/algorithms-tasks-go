package fibfrog

import (
	"fmt"
	"sort"
)

var cacheResults = make(map[string]int)

func Solution(A []int) int {
	possibleStack := NewStack()
	for idx := 0; idx <= len(A); idx++ {
		// add possibles until finish
		stepNumber := idx + 1
		if isItFibonachi(stepNumber) && idx != len(A) && A[idx] == 1 {
			possibleStack.push(idx)
		}

		if idx == len(A) { // Is it final?
			if isItFibonachi(stepNumber) {
				return 1
			} else {
				if possibleStack.isEmpty() {
					return -1
				}
				var results []int
				for !possibleStack.isEmpty() {
					tailOfArray := A[possibleStack.pop()+1:] // first inclusive, second exclusive
					keyToSearch := arrayToString(tailOfArray)
					cachedResult := cacheResults[keyToSearch]
					if cachedResult != 0 {
						if cachedResult != -1 {
							results = append(results, cachedResult)
						}
					} else {
						result := Solution(tailOfArray)
						cacheResults[keyToSearch] = result
						if result != -1 {
							results = append(results, result)
						}
					}
				}
				if len(results) == 0 {
					return -1
				} else {
					sort.Ints(results)
					return results[0] + 1
				}
			}

		}
	}

	return -1
}

func arrayToString(arr []int) string {
	return fmt.Sprint(arr)
}

func isItFibonachi(value int) bool {
	fibonachi := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393}
	return binarySearch(value, fibonachi)
}

func binarySearch(needle int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != needle {
		return false
	}

	return true
}

type Stack struct {
	data []int
	idx  int
}

func NewStack() *Stack {
	return &Stack{
		make([]int, 100000),
		-1,
	}
}

func (s *Stack) pop() int {
	if s.idx == -1 {
		return -1
	}
	value := s.data[s.idx]
	s.idx = s.idx - 1
	return value
}

func (s *Stack) push(value int) {
	s.idx = s.idx + 1
	s.data[s.idx] = value
}

func (s *Stack) isEmpty() bool {
	return s.idx == -1
}
