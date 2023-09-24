package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"pushswap/funcs"
)

func main() {
	funcs.InitializeStackA()
	if funcs.IsSorted(funcs.StackA, funcs.StackB) {
		fmt.Println("input alrady Sorted")
		return
	}

	file, err := os.Create("Result-log.txt")
	if err != nil {
		fmt.Println("Error: while Createing Result-log.txt file")
		return
	}
	defer file.Close()
	numsToStr := fmt.Sprint(funcs.StackA)

	file.WriteString("the input:\n\"" + numsToStr + "\"\n")

	instructions := ""
	funcs.UInts(funcs.StackA) // replace nums with there index

	StackaToStr := fmt.Sprint(funcs.StackA)
	file.WriteString("stake A:\n\"" + StackaToStr + "\"\n")

	instructions = funcs.Sort(funcs.StackA)
	sortedStackaToStr := fmt.Sprint(funcs.StackA)
	file.WriteString("sorted stake A:\n\"" + sortedStackaToStr + "\"\n")

	file.WriteString("instructions:\n\"\n" + instructions + "\"\n")
	file.WriteString("steps: " + strconv.Itoa(len(strings.Split(instructions, "\n"))-1))

	// fmt.Println("steps: ", len(strings.Split(instructions, "\n"))-1)
	fmt.Println(instructions)
}
