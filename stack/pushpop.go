package stack

import "fmt"

type (
	S []int64
)

// IsEmpty: check if stack is empty
func (s *S) AIsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *S) Push(n int64) {
	*s = append(*s, n) // Simply append the new value to the end of the array
	// the end of the array will be the top of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *S) Pop() (int64, bool) {
	if s.AIsEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *S) PopStack() {
	for len(*s) > 0 {
		x, success := s.Pop()
		if success {
			fmt.Println(x)
		}
	}
}
