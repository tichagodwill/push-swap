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

// InitializeStackA initializes stack A with integers provided as command-line arguments.
// It expects a single argument, a space-separated string of integers.
// If no argument is provided or more than one argument is given, the program exits.
func InitializeStackA() {
	if len(os.Args) == 1 || os.Args[1] == "" {
		// If no argument or an empty argument is provided, exit gracefully.
		os.Exit(0)
	} else if len(os.Args) != 2 {
		// If more than one argument is provided, display usage and exit.
		fmt.Println("Usage: ./push-swap Replace [int1 int2 int3...] with the integers you want to sort.")
		fmt.Println("Example: ./push-swap \"2 1 3 6 5 8\"")
		os.Exit(0)
	}
	PushToStack(os.Args[1])
}

// PushToStack parses a space-separated string of integers and pushes them onto stack A.
// It checks for duplicate values and reports an error if duplicates are found.
func PushToStack(str string) {
	numStr := strings.Fields(str)
	// Push from the last element to the first so the first will be on the top of the stack.
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

// CheckDuplicate checks if an integer value n already exists in stack A.
// It returns true if a duplicate is found, otherwise returns false.
func CheckDuplicate(n int, StackA stack.S) bool {
	for i := 0; i < len(StackA); i++ {
		if n == StackA[i] {
			return true
		}
	}
	return false
}
