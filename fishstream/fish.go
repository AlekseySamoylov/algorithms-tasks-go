package fishstream





func Solution(A []int, B []int) int {
	size := A
	direction := B
	upstream := 0
	downstream := 1
	currentPredatorIdx := -1
	currentPredatorSize := -1
	upstreamDirectionEaten := -1
	downstreamDirectionEated := -2
	for fishIdx := 0; fishIdx < len(size); fishIdx++ {
		if direction[fishIdx] == upstream {
			if currentPredatorIdx != -1 {
				if currentPredatorSize > size[fishIdx] {
					direction[fishIdx] = upstreamDirectionEaten
				} else {
					currentPredatorSize = -1
					direction[currentPredatorIdx] = downstreamDirectionEated
				}
			} else {
				continue
			}
		} else if direction[fishIdx] == downstream {
			if currentPredatorIdx != -1 {
				if currentPredatorSize < size[fishIdx] {
					currentPredatorIdx = fishIdx
					currentPredatorSize = size[fishIdx]
				}
			} else {
				currentPredatorIdx = fishIdx
				currentPredatorSize = size[fishIdx]
			}
		}
	}
	currentPredatorIdx = -1
	currentPredatorSize = -1
	for fishIdx := len(size) - 1; fishIdx > 0; fishIdx-- {
		if direction[fishIdx] == upstream || direction[fishIdx] == -1 {
			if currentPredatorIdx != -1 {
				if currentPredatorSize < size[fishIdx] {
					currentPredatorIdx = fishIdx
					currentPredatorSize = size[fishIdx]
				}
			} else {
				currentPredatorIdx = fishIdx
				currentPredatorSize = size[fishIdx]
			}
		} else if direction[fishIdx] == downstream || direction[fishIdx] == -2 {
			if currentPredatorIdx != -1 {
				if currentPredatorSize > size[fishIdx] {
					direction[fishIdx] = downstreamDirectionEated
				} else {
					currentPredatorSize = -1
					direction[currentPredatorIdx] = upstreamDirectionEaten
				}
			} else {
				continue
			}
		}
	}
	aliveFishNumber := 0
	for fishIdx := 0; fishIdx < len(size); fishIdx++ {
		if direction[fishIdx] != upstreamDirectionEaten && direction[fishIdx] != downstreamDirectionEated {
			aliveFishNumber++
		}
	}
	return aliveFishNumber
}
