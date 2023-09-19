package funcs

import "pushswap/stack"

func PB(sA *stack.S, sB *stack.S) {
	a, err := (*stack.S)(sA).Pop()
	if err {
		*sB = append(*sB, a)
	}
}

func PA(sA *stack.S, sB *stack.S) {
	b, err := (*stack.S)(sB).Pop()
	if err {
		*sA = append(*sA, b)
	}
}
