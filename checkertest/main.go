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
	// Before starting we should initialize the stack
	funcs.InitializeStackA()

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
				fmt.Printf("Error: cannot apply the instruction %v\n", input)
				fmt.Println("KO")
				os.Exit(1)
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
