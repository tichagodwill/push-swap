package funcs

func PB() bool {
	a, canPop := StackA.Pop()
	if canPop {
		StackB = append(StackB, a)
		return true
	}
	return false
}

func PA() bool {
	b, canPop := StackB.Pop()
	if canPop {
		StackA = append(StackA, b)
		return true
	}
	return false
}
