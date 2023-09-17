package main

import (
	"fmt"

	"pushswap/funcs"
)

func main() {
	var stack funcs.Stack // create a stack variable of type Stack
	stack.Push("this")
	stack.Push("is")
	stack.Push("sparta!!")

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y {
			fmt.Println(x)
		}
	}
}
