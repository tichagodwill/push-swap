package funcs

func RA() bool {
	len := len(StackA)
	if len < 2 {
		return false
	}
	firstElement := StackA[len-1]
	for i := len - 1; i > 0; i-- {
		StackA[i] = StackA[i-1]
	}
	StackA[0] = firstElement
	return true
}

func RB() bool {
	len := len(StackB)
	if len < 2 {
		return false
	}
	firstElement := StackB[len-1]
	for i := len - 1; i > 0; i-- {
		StackB[i] = StackB[i-1]
	}
	StackB[0] = firstElement
	return true
}

func RR() bool {
	if RA() && RB() {
		return true
	}
	return false
}
