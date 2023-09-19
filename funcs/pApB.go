package funcs

import "pushswap/stack"

type (
	StackA stack.S // []int64
	StackB stack.S // []int64
)

func PB(sA *StackA, sB *StackB) {
	a, err := (*stack.S)(sA).Pop()
	if !err {
		*sB = append(*sB, a)
	}
}

func PA(sA *StackA, sB *StackB) {
	b, err := (*stack.S)(sB).Pop()
	if !err {
		*sA = append(*sA, b)
	}
}
