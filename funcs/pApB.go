package funcs

import "pushswap/stack"

func PB(sA *stack.S, sB *stack.S) {
	a, canPop := (*stack.S)(sA).Pop()
	if canPop {
		*sB = append(*sB, a)
	}
}

func PA(sA *stack.S, sB *stack.S) {
	b, canPop := (*stack.S)(sB).Pop()
	if canPop {
		*sA = append(*sA, b)
	}
}
