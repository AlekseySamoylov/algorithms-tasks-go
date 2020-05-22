package ctplslises




func Solution(M int, A []int) int {
	counter := 0
	duplicateArray := make([]int, M + 1)
	if M == 0 {
		return len(A)
	}
	for tailIdx := 0; tailIdx < len(A); tailIdx++ {
		for headIdx := tailIdx; headIdx < len(A); headIdx++ {
			if headIdx < tailIdx {
				continue
			} else if counter > 1000000000 {
				return 1000000000
			} else if tailIdx == headIdx {
				counter++
				duplicateArray[A[headIdx]] = 1
			} else if duplicateArray[A[headIdx]] != 0 {
				duplicateArray = make([]int, M + 1)
				break
			} else {
				counter++
				duplicateArray[A[headIdx]] = 1
			}
		}
		duplicateArray = make([]int, M + 1)
	}
	return counter
}

func Solution_2(M int, A []int) int {
	counter := 0
	tempCounter := 0
	duplicateArray := make([]int, M + 1)
	if M == 0 {
		return len(A)
	}
	for headIdx := 0; headIdx < len(A); headIdx++ {
		value := A[headIdx]
		if counter > 1000000000 {
			return 1000000000
		} else if duplicateArray[value] != 0 {
			duplicateArray = make([]int, M+1)
			counter += tempCounter * (tempCounter + 1)/2
			duplicateArray[value] = 1
			tempCounter = 1

		} else {
			duplicateArray[value] = 1
			tempCounter++
		}

	}
	counter += tempCounter * (tempCounter + 1)/2

	return counter
}


func Solution_1(M int, A []int) int {
	markerArray := make([]int, M)
	resultMap := make(map[int]int)

	counter := 0
	periodCounter := 0
	periodSliceCounter := 0
	sliceLangth := 1
	for _, element := range A {
		if markerArray[element - 1] == 0 {
			markerArray[element - 1] = 1
			sliceLangth++
			periodCounter++
		} else {
			markerArray = make([]int, M)
			result, isFound := resultMap[periodCounter]
			if isFound {
				periodSliceCounter = result
			} else {
				for i := 1; i <= periodCounter; i++ {
					for j := 1; j <= i; j++ {
						periodSliceCounter++
					}
				}
			}
			resultMap[periodCounter] = periodSliceCounter
			periodCounter = 0
			counter += periodSliceCounter
			periodSliceCounter = 0
			markerArray = make([]int, M)
			markerArray[element - 1] = 1
			sliceLangth++
			periodCounter++
		}
	}
	markerArray = make([]int, M)
	result, isFound := resultMap[periodCounter]
	if isFound {
		periodSliceCounter = result
	} else {
		for i := 1; i <= periodCounter; i++ {
			for j := 1; j <= i; j++ {
				periodSliceCounter++
			}
		}
	}
	periodCounter = 0
	counter += periodSliceCounter
	return counter
}
