package funcs

import "fmt"

// validInstructions is a map that associates each instruction string with its corresponding function.
var validInstructions = map[string]func() bool{
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

// IsInstruction checks if a given string is a valid instruction.
func IsInstruction(str string) bool {
	return validInstructions[str] != nil
}

// ApplyInstruction applies the given instruction if it's valid and returns true.
// If the instruction is not valid, it prints an error message and returns false.
func ApplyInsruction(str string) bool {
	if !IsInstruction(str) {
		fmt.Printf("Cannot Apply %v instruction\n", str) 
		return false
	}
	return validInstructions[str]()
}
