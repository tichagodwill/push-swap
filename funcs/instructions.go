package funcs

import "fmt"

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

func IsInstruction(str string) bool {
	return validInstructions[str] != nil
}

func ApplyInsruction(str string) bool {
	if !IsInstruction(str) {
		fmt.Printf("Cannot Apply %v instruction\n", str) 
		return false
	}
	return validInstructions[str]()
}
