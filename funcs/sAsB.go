package funcs

//Swap the first two elements of stack A.
func SA() bool {
	len := len(StackA)
	if len < 2 {
		return false
	}
	StackA[len-2], StackA[len-1] = StackA[len-1], StackA[len-2]
	return true
}

//Swap the first two elements of stack B.
func SB() bool {
	len := len(StackB)
	if len < 2 {
		return false
	}
	StackB[len-2], StackB[len-1] = StackB[len-1], StackB[len-2]
	return true
}

//Execute `sa` and `sb` simultaneously.
func SS() bool {
	if SA() && SB() {
		return true
	}
	return false
}
