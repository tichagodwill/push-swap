package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"pushswap/funcs"
	"pushswap/stack"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:")
		return
	}
	var sa stack.S // create a stack variable of type Stack
	var sb stack.S // create a stack variable of type Stack
	PushArgsToStack(os.Args[1], &sa)
	funcs.PA(&sa, &sb)
	PopStack(&sa)
}

func PushArgsToStack(str string, stack *stack.S) {
	numStr := strings.Fields(os.Args[1])
	// Push from the last element to the first so the first will be on the top of the stack
	for i := len(numStr) - 1; i >= 0; i-- {
		n, err := strconv.ParseInt(numStr[i], 10, 64)
		if err != nil {
			fmt.Println("Error Parsing integer ", err)
			return
		}
		stack.Push(n)
	}
}

func PopStack(stack *stack.S) {
	for len(*stack) > 0 {
		x, y := stack.Pop()
		if y {
			fmt.Println(x)
		}
	}
}
