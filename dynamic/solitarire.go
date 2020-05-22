package dynamic

func Solution(A []int) int {
	startIndexToBestValueIndexMap := make(map[int]int)
	initialInt := 0 - 10001
	sum0 := 0

	for startIndex := 0; startIndex < len(A); startIndex++ {
		bestValue := initialInt
		for bestValueIndex := startIndex + 1; bestValueIndex <= (startIndex+7) && bestValueIndex < len(A); bestValueIndex++ {

			if bestValue < A[bestValueIndex] {
				startIndexToBestValueIndexMap[startIndex] = bestValueIndex
				bestValue = A[bestValueIndex]

				if bestValueIndex >= len(A)-7 {
					println(">>")
					println(bestValueIndex)
					if sum0 < bestValue {
						sum0 = bestValue
					}
				}
			}
		}
	}

	doIt := true
	nextIndex := 0
	sum := 0
	for doIt {
		sum += A[nextIndex]
		nextIndex, doIt = startIndexToBestValueIndexMap[nextIndex]
	}

	println(startIndexToBestValueIndexMap)
	for key, value := range startIndexToBestValueIndexMap {
		println(key)
		println(value)
		println("******")
	}
	return sum
}

func SolutionOld(A []int) int {
	initialInt := 0 - 10001
	maxStep := 6
	highestValueInPeriod := initialInt
	highestValueInPeriodIdx := 1
	result := A[0]
	nextStepIdx := 1

	for i := 1; i < len(A); i++ {
		for internalStep := i; internalStep < maxStep; internalStep++ {
			if internalStep == len(A)-1 {
				if highestValueInPeriod > A[internalStep] && (highestValueInPeriod+A[internalStep]) > A[internalStep] {
					nextStepIdx = highestValueInPeriodIdx
					break
				}
				highestValueInPeriod = A[internalStep]
				nextStepIdx = internalStep
				break
			}
			if highestValueInPeriod == initialInt {
				highestValueInPeriod = A[internalStep]
				nextStepIdx = internalStep + 1
				highestValueInPeriodIdx = internalStep
				continue
			}
			if A[internalStep] > highestValueInPeriod {
				highestValueInPeriod = A[internalStep]
				nextStepIdx = internalStep + 1
				highestValueInPeriodIdx = internalStep
			}
		}
		result += highestValueInPeriod
		highestValueInPeriod = initialInt
		i = nextStepIdx
		highestValueInPeriodIdx = i
		if i == len(A) {
			result += A[i]
		}
	}
	return result
}
