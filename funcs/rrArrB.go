package funcs

//Reverse rotate stack A (shift down all elements by 1; the last element becomes the first).
func RRA() bool {
	len := len(StackA)
	if len < 2 {
		return false
	}
	lastElement := StackA[0]
	for i := 0; i < len-1; i++ {
		StackA[i] = StackA[i+1]
	}
	StackA[len-1] = lastElement
	return true
}

//Reverse rotate stack B.
func RRB() bool {
	len := len(StackB)
	if len < 2 {
		return false
	}
	lastElement := StackB[0]
	for i := 0; i < len-1; i++ {
		StackB[i] = StackB[i+1]
	}
	StackB[len-1] = lastElement
	return true
}

//Execute `rra` and `rrb` simultaneously.
func RRR() bool {
	if RRA() && RRB() {
		return true
	}
	return false
}
