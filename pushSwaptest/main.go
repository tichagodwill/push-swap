package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"pushswap/funcs"
)

func main() {
	// Initialize stack A with command-line arguments.
	funcs.InitializeStackA()

	// Check if stack A is already sorted; if so, print a message and exit.
	if funcs.IsSorted(funcs.StackA, funcs.StackB) {
		fmt.Println("Input already sorted")
		return
	}

	// Create a result log file.
	file, err := os.Create("Result-log.txt")
	if err != nil {
		fmt.Println("Error: while creating Result-log.txt file")
		return
	}
	defer file.Close() // Ensure the file is closed when done.

	// Convert stack A to a string and write it to the log file.
	numsToStr := fmt.Sprint(funcs.StackA)
	file.WriteString("the input:\n\"" + numsToStr + "\"\n")

	// Perform some operations on stack A (e.g., replace nums with their index).
	funcs.UInts(funcs.StackA)

	// Convert stack A to a string and write it to the log file.
	StackaToStr := fmt.Sprint(funcs.StackA)
	file.WriteString("stack A:\n\"" + StackaToStr + "\"\n")

	// Sort stack A and record the instructions.
	instructions := funcs.Sort(funcs.StackA)

	// Convert the sorted stack A to a string and write it to the log file.
	sortedStackaToStr := fmt.Sprint(funcs.StackA)
	file.WriteString("sorted stack A:\n\"" + sortedStackaToStr + "\"\n")

	// Write the sorting instructions and the number of steps to the log file.
	file.WriteString("instructions:\n\"\n" + instructions + "\"\n")
	file.WriteString("steps: " + strconv.Itoa(len(strings.Split(instructions, "\n"))-1))

	// Print the instructions to the standard output.
	fmt.Println(instructions)
}
