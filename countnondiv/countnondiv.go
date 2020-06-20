package countnondiv

import "math/rand"

// https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_non_divisible/
func Solution(A []int) []int {
	aSize := len(A)
	maxValue := A[0]
	for idx := 0; idx < aSize; idx++ {
		element := A[idx]
		if maxValue < element {
			maxValue = element
		}
	}

	remainDividersCountList := make([]int, maxValue+1)

	for idx := 0; idx < aSize; idx++ {
		element := A[idx]
		if remainDividersCountList[element] == 0 {
			remainDividersCountList[element] = aSize - 1
		} else {
			remainDividersCountList[element] = remainDividersCountList[element] - 1
		}
	}
	cachedResultMap := make([]int, maxValue+1)
	alreadyCalculated := make([]int, maxValue+1)
	alreadyCalculatedDuplicated := make([]int, maxValue+1)
	caluclatedMap := make(map[int][]int)
	for idx := 0; idx < aSize; idx++ {
		element := A[idx]
		if alreadyCalculated[element] == 0 {
			for multiplier := 2; multiplier <= maxValue/element; multiplier++ {
				multResult := element * multiplier
				if multResult > maxValue {
					break
				} else {
					cachedResult := cachedResultMap[multResult]
					if cachedResult > 0 {
						cachedResultMap[multResult] = cachedResult + 1
					} else {
						cachedResultMap[multResult] = 1
					}
					caluclatedMap[element] = append(caluclatedMap[element], multResult)
				}
			}
			alreadyCalculated[element] = 1
		} else if alreadyCalculatedDuplicated[element] == 0 {
			multiplier := aSize - (remainDividersCountList[element] + 1)
			list := caluclatedMap[element]
			for repIdx := 0; repIdx < len(list); repIdx++ {
				repElem := list[repIdx]
				cachedResultMap[repElem] = cachedResultMap[repElem] + (1 * multiplier)
			}
			alreadyCalculatedDuplicated[element] = 1
		}
	}

	result := make([]int, aSize)
	for idx := 0; idx < aSize; idx++ {
		element := A[idx]
		result[idx] = remainDividersCountList[element] - cachedResultMap[element]
	}
	return result
}

func fasterSolution(A []int) []int {
	aSize := len(A)
	maxValue := A[0]
	for idx := 0; idx < aSize; idx++ {
		element := A[idx]
		if maxValue < element {
			maxValue = element
		}
	}

	remainDividersCountList := make([]int, maxValue+1)

	for idx := 0; idx < aSize; idx++ {
		element := A[idx]
		if remainDividersCountList[element] == 0 {
			remainDividersCountList[element] = aSize - 1
		} else {
			remainDividersCountList[element] = remainDividersCountList[element] - 1
		}
	}
	cachedResultMap := make([]int, maxValue+1)
	for idx := 0; idx < aSize; idx++ {
		element := A[idx]
		for multiplier := 2; multiplier <= maxValue; multiplier++ {
			multResult := element * multiplier
			if multResult > maxValue {
				break
			} else {
				cachedResult := cachedResultMap[multResult]
				if cachedResult > 0 {
					cachedResultMap[multResult] = cachedResult + 1
				} else {
					cachedResultMap[multResult] = 1
				}
			}
		}
	}

	result := make([]int, aSize)
	for idx := 0; idx < aSize; idx++ {
		element := A[idx]
		result[idx] = remainDividersCountList[element] - cachedResultMap[element]
	}
	return result
}

func BitFasterSolution(A []int) []int {
	//maxValue := 100000
	sizeA := len(A)
	maxSize := 0
	for outerIdx := 0; outerIdx < len(A); outerIdx++ {
		outerElement := A[outerIdx]
		if outerElement > maxSize {
			maxSize = outerElement
		}
	}
	cachedResultMap := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		for outerIdx := 0; outerIdx < sizeA; outerIdx++ {
			outerElement := A[outerIdx]
			if (i+1)%outerElement != 0 {
				cachedResultMap[i] = cachedResultMap[i] + 1
			}
		}

	}
	result := make([]int, sizeA)
	for outerIdx := 0; outerIdx < len(A); outerIdx++ {
		result[outerIdx] = cachedResultMap[A[outerIdx]-1]
	}

	return result
}

func SolutionSlow(A []int) []int {
	result := make([]int, len(A))

	for outerIdx := 0; outerIdx < len(A); outerIdx++ {
		result[outerIdx] = 0
		outerElement := A[outerIdx]
		for innerIdx := 0; innerIdx < len(A); innerIdx++ {
			if outerIdx == innerIdx {
				continue
			}
			innerElement := A[innerIdx]

			if outerElement%innerElement == 0 {
			} else {
				result[outerIdx] = result[outerIdx] + 1
			}
		}
	}

	return result
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
