package funcs

import "fmt"

func RRA() bool {
	len := len(StackA)
	if len < 2 {
		fmt.Println("Error RRA")
		return false
	}
	lastElement := StackA[0]
	for i := 0; i<len-1; i++ {
		StackA[i] = StackA[i+1]
	}
	StackA[len-1] = lastElement
	return true
}

func RRB() bool {
	len := len(StackB)
	if len < 2 {
		fmt.Println("Error RRB")
		return false
	}
	lastElement := StackB[0]
	for i := 0; i<len-1; i++ {
		StackB[i] = StackB[i+1]
	}
	StackB[len-1] = lastElement
	return true
}

func RRR() bool {
	if RRA() && RRB() {
		return true
	}
	fmt.Println("Error RRR")
	return false
}