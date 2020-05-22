package greedy



func Solution(M int, A []int) int {
	max := 0
	currentRope := 0
	for tail := 0; tail < len(A); tail++ {
		if A[tail] >= M {
			max++
			currentRope = 0
			continue
		}
		currentRope += A[tail]
		if (currentRope) >= M {
			max++
			currentRope = 0
			continue
		}
	}
	return max
}
func Solution2(M int, A []int) int {
	max := 0
	currentRope := 0
	for tail := 0; tail < len(A); tail++ {
		if A[tail] >= M {
			max++
			continue
		}
		currentRope = A[tail]
		for head := tail + 1; head < len(A); head++ {
			currentRope = currentRope + A[head]
			if currentRope >= M {
				max++
				tail = head
				currentRope = 0
				break
			}
		}
	}
	return max
}

func Solution_old(M int, A []int) int {
	max := 0
	currentAttemptRopes := 1
	lengthReached := false
	for tail := 0; tail < len(A); tail++ {
		if A[tail] >= M {
			if max < currentAttemptRopes {
				max = currentAttemptRopes
			}
			continue
		}
		for head := tail + 1; head < len(A); head++ {
			currentAttemptRopes++
			currentRope := A[tail] + A[head]
			if currentRope >= M {
				lengthReached = true
				if max < currentAttemptRopes {
					max = currentAttemptRopes
				}
			}
			if lengthReached && tail < head {
				tail++
				lengthReached = false
				currentAttemptRopes--
			}
		}
		currentAttemptRopes = 0
		lengthReached = false
	}
	return max
}
