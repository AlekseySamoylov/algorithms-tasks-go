package nestedstring

func Solution(A string) bool {

	current := 0
	open := '('
	closed := ')'
	for _, char := range A {
		if char == open {
			current++
		} else if char == closed {
			current--
		}
		if current < 0 {
			return false
		}
	}

	return current == 0
}
