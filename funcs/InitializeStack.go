package funcs

import (
	"fmt"
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

func InitializeStackA() {
	if len(os.Args) == 1 || os.Args[1] == ""{
		os.Exit(0)
	} else if len(os.Args) != 2 {
		fmt.Println("usage:")
		os.Exit(0)
	}
	PushToStack(os.Args[1])
}

func PushToStack(str string) {
	numStr := strings.Fields(str)
	// Push from the last element to the first so the first will be on the top of the stack
	for i := len(numStr) - 1; i >= 0; i-- {
		n, err := strconv.ParseInt(numStr[i], 10, 64)
		if err != nil {
			log.Fatal("Error Parsing integer ", err)
		}
		if CheckDuplicate(int(n), StackA) {
			log.Fatalf("Error: unexpected duplicate value of num:%d detected.\n", n)
		}
		StackA.Push(int(n))
	}
}

func CheckDuplicate(n int, StackA stack.S) bool {
	for i := 0; i < len(StackA); i++ {
		if n == StackA[i] {
			return true
		}
	}
	return false
}
