package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"pushswap/funcs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage:")
		return
	}
	var stacka funcs.StackA // create a stack variable of type Stack
	numStr := strings.Fields(os.Args[1])
	// Push from the last element to the first so the first will be on the top of the stack
	for i := len(numStr) - 1; i >= 0; i-- {
		n, err := strconv.ParseInt(numStr[i], 10, 64)
		if err != nil {
			fmt.Println("Error Parsing integer ", err)
			return
		}
		stacka.Push(n)
	}
	for len(stacka) > 0 {
		x, y := stacka.Pop()
		if y {
			fmt.Println(x)
		}
	}
}
