package funcs

import "fmt"

func RA() bool {
	len := len(StackA)
	if len < 2 {
		fmt.Println("Error RA")
		return false
	}
	firstElement := StackA[len-1]
	for i := len-1; i>0; i-- {
		StackA[i] = StackA[i-1]
	}
	StackA[0] = firstElement
	return true
}

func RB() bool {
	len := len(StackB)
	if len < 2 {
		fmt.Println("Error RB")
		return false
	}
	firstElement := StackB[len-1]
	for i := len-1; i>0; i-- {
		StackB[i] = StackB[i-1]
	}
	StackB[0] = firstElement
	return true
}

func RR() bool {
	if RA() && RB() {
		return true
	}
	fmt.Println("Error RR")
	return false
}