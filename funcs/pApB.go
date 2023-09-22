package funcs

import "fmt"

func PB() bool {
	a, canPop := StackA.Pop()
	if canPop {
		StackB = append(StackB, a)
		return true
	}
	fmt.Println("Error PB")
	return false
}

func PA() bool {
	b, canPop := StackB.Pop()
	if canPop {
		StackA = append(StackA, b)
		return true
	}
	fmt.Println("Error PA")
	return false
}
