package stack

import "fmt"

type (
	S []int
)

// IsEmpty: check if stack is empty
func (s *S) AIsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *S) Push(n int) {
	*s = append(*s, n) // Simply append the new value to the end of the array
	// the end of the array will be the top of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *S) Pop() (int, bool) {
	if s.AIsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// PopStack pops all elements from the stack and prints them.
func (s *S) PopStack() {
	// Continue popping elements from the stack until it's empty.
	for len(*s) > 0 {
		// Pop an element from the stack and check if it was successful.
		x, success := s.Pop()

		// If the pop operation was successful, print the popped element.
		if success {
			fmt.Println(x)
		}
	}
}
