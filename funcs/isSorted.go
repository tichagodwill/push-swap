package funcs

import "pushswap/stack"

func IsSorted(sa stack.S, sb stack.S) bool {
	if len(sb) != 0 {
		return false
	}
	prev, _ := sa.Pop()
	for len(sa) > 0 {
		x, success := sa.Pop()
		if !success {
			break
		}
		if x < prev {
			return false
		}
		prev = x
	}
	return true
}
