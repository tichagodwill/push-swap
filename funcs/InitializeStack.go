package funcs

import (
	"log"
	"os"
	"strconv"
	"strings"

	"pushswap/stack"
)

var (
	StackA stack.S
	StackB stack.S
)

func InitializeStackA(str string) {
	PushArgsToStack(str)
}

func PushArgsToStack(str string) {
	numStr := strings.Fields(os.Args[1])
	// Push from the last element to the first so the first will be on the top of the stack
	for i := len(numStr) - 1; i >= 0; i-- {
		n, err := strconv.ParseInt(numStr[i], 10, 64)
		if err != nil {
			log.Fatal("Error Parsing integer ", err)
		}
		StackA.Push(n)
	}
}
