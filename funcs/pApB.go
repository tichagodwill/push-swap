package funcs

//Push the top element of stack A to stack B.
func PB() bool {
	a, canPop := StackA.Pop()
	if canPop {
		StackB = append(StackB, a)
		return true
	}
	return false
}

//Push the top element of stack B to stack A.
func PA() bool {
	b, canPop := StackB.Pop()
	if canPop {
		StackA = append(StackA, b)
		return true
	}
	return false
}
