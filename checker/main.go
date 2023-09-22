package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"pushswap/funcs"
)

func main() {
	if len(os.Args) == 1 {
		return
	} else if len(os.Args) != 2 {
		fmt.Println("usage:")
		return
	}
	// Before starting we should initialize the stack
	funcs.InitializeStackA(os.Args[1])

	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Read input from standard input
	for scanner.Scan() {
		// Retrieve the input from the scanner
		input := scanner.Text()
		input = strings.ToLower(input)
		// Check if the input is a valid instruction
		if input == "" {
			break
		} else if funcs.IsInstruction(input) {
			if !funcs.ApplyInsruction(input) {
				log.Fatal("Error: cannot apply the instruction")
			}
		} else {
			log.Fatal("Error: invalid instruction")
		}
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	if funcs.IsSorted(funcs.StackA, funcs.StackB) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
