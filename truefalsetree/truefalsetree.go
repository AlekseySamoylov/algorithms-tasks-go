package truefalsetree

func Solution(A []bool) int {
	totalHeight := len(A)
	totalNodes := (totalHeight * (totalHeight + 1)) / 2
	totalEmptyNodes := 0
	var emptyBasesList = make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if !A[i] {
			totalEmptyNodes++
		} else {
			emptyBasesList[i] = totalEmptyNodes
			totalEmptyNodes = 0
		}
	}
	if totalEmptyNodes > 0 {
		emptyBasesList[len(A)-1] = totalEmptyNodes
	}
	trueNodes := totalNodes
	for i := 0; i < len(emptyBasesList); i++ {
		if emptyBasesList[i] > 0 {
			falseTreeHeight := emptyBasesList[i]
			falseNodes := (falseTreeHeight * (falseTreeHeight + 1)) / 2
			trueNodes = trueNodes - falseNodes
		}
	}
	if trueNodes >= 1000000000 {
		return 1000000000
	}
	return trueNodes
}

func Solution2(A []bool) int {
	boolCount := 0
	for i := 0; i < len(A); i++ {
		if A[i] {
			boolCount++
			if boolCount >= 1000000000 {
				return 1000000000
			}
		}
	}
	currentArray := A
	for len(currentArray) > 1 {
		nextArray := make([]bool, len(currentArray)-1)
		for i := 0; i+1 < len(currentArray); i++ {
			aNode := currentArray[i]
			bNode := currentArray[i+1]
			if aNode || bNode {
				boolCount++
				if boolCount >= 1000000000 {
					return 1000000000
				}
				nextArray[i] = true
			} else {
				nextArray[i] = false
			}
		}
		currentArray = nextArray
	}

	return boolCount
}
