package main

import (
	"fmt"
	"os"

	"pushswap/funcs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:")
		return
	}
	funcs.InitializeStackA(os.Args[1])
	funcs.StackA.PopStack()
	// var sa stack.S // create a stack variable of type Stack
	// var sb stack.S // create a stack variable of type Stack
	// funcs.PB(&sa, &sb)
	// funcs.SA(&sa)
	// PopStack(sa)
	// PopStack(sb)
}
