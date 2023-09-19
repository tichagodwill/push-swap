package funcs

import "pushswap/stack"

func PB(sA *stack.S, sB *stack.S) bool {
	a, canPop := sA.Pop()
	if canPop {
		*sB = append(*sB, a)
		return true
	}
	return false
}

func PA(sA *stack.S, sB *stack.S) bool {
	b, canPop := sB.Pop()
	if canPop {
		*sA = append(*sA, b)
		return true
	}
	return false
}
