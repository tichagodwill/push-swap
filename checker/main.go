package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"pushswap/funcs"
	"pushswap/stack"
)

var validInstructions = map[string]func() bool{
	"pa":  funcs.PA,
	"pb":  funcs.PB,
	"sa":  funcs.SA,
	"sb":  funcs.SB,
	"ss":  funcs.SS,
	"ra":  funcs.RA,
	"rb":  funcs.RB,
	"rr":  funcs.RR,
	"rra": funcs.RRA,
	"rrb": funcs.RRB,
	"rrr": funcs.RRR,
}

func main() {
	if len(os.Args) == 1 {
		return
	} else if len(os.Args) != 2 {
		fmt.Println("usage:")
		return
	}
	funcs.InitializeStackA(os.Args[1])

	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Read input from standard input
	for scanner.Scan() {
		// Retrieve the input from the scanner
		input := scanner.Text()
		// Check if the input is a valid instruction
		if input == "" {
			break
		} else if isInstruction(input) {
			if !applyInsruction(input) {
				log.Fatal("Error: cannot apply the instruction")
			}
		} else {
			log.Fatal("Error: invalid instruction")
		}
	}

	if isSorted(funcs.StackA, funcs.StackB) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}

func isInstruction(str string) bool {
	str = strings.ToLower(str)
	return validInstructions[str] != nil
}

func applyInsruction(str string) bool {
	return validInstructions[str]()
}

func isSorted(sa stack.S, sb stack.S) bool {
	if len(sb) != 0 {
		return false
	}
	prev, _ := sa.Pop()
	for len(sa) > 0 {
		x, success := sa.Pop()
		if !success {
			break
		}
		if x < prev {
			return false
		}
		prev = x
	}
	return true
}
