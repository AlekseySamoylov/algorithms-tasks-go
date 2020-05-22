package smallest

func Solution(A int, B int, K int) int {
	ifZero := 0
	if A == 0 {
		ifZero = 1
	}
	if K != 1 && K > B {
		return ifZero
	}
	if K == 1 {
		return (B - A) + 1
	}
	count := 0
	temp := (A / K) * K
	for temp <= B {
		if temp >= A && temp <= B {
			count++
		}
		temp += K
	}

	return count
}

func Solution3Wrong(A int, B int, K int) int {
	count := 0
	if A == 0 && B < K{
		return 1
	}
	if K > B && K != 1 {
		return 0
	}
	if A <= K && K != 1 {
		return B / K
	}
	if B - A >= K {
		return (1 + B - A) / K
	}
	temp := B
	cont := true
	for cont {
		temp = temp / K
		if temp <= A || temp == 1 {
			cont = false
			return count
		}
		count++
	}
	return 0
}

func Solution2(A int, B int, K int) int {
	// 2,3,4,5
	// 5,6,7,8
	count := 0
	for i := A; i <= B; i++ {
		if i % K == 0 {
			count++
		}
	}

	return count
}
