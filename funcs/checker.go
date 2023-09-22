package funcs

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Checker() {
	instructions := []string{}

	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Read input from standard input
	for scanner.Scan() {
		// Retrieve the input from the scanner
		input := scanner.Text()
		// Check if the input is a valid instruction
		if isInstruction(input) {
			instructions = append(instructions, input)
		} else {
			log.Fatal("Error: invalid instruction")
		}
		// if len(instructions) == 7 {
		// 	break
		// }
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	// Apply the instructions on the stack
	// ApplyInsructions(instructions)
}

func isInstruction(str string) bool  {
	str = strings.ToLower(str)
	validInstructions := map[string]func()bool{
		"pa":  PA,
		"pb":  PB,
		"sa":  SA,
		"sb":  SB,
		"ss":  SS,
		"ra":  RA,
		"rb":  RB,
		"rr":  RR,
		"rra": RRA,
		"rrb": RRB,
		"rrr": RRR,
	}
	if validInstructions[str] != nil {
		return validInstructions[str]()
	}
	return false
}

// func ApplyInsructions(inst []string) bool {

// }

