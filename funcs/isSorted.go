package funcs

import "pushswap/stack"

// IsSorted checks if the given stack `sa` is sorted in ascending order.
// It returns true if the stack is sorted, false otherwise.
// It also checks if the stack `sb` is empty since Push-Swap requires stack `sb` to be empty for a valid sorting result.
func IsSorted(sa stack.S, sb stack.S) bool {
	// Check if stack `sb` is empty; it should be empty for a valid sorting result.
	if len(sb) != 0 {
		return false
	}

	// Initialize the previous element with the top element of stack `sa`.
	prev, _ := sa.Pop()

	// Iterate through the remaining elements of stack `sa`.
	for len(sa) > 0 {
		x, success := sa.Pop()
		if !success {
			break
		}
		// Compare the current element `x` with the previous element.
		// If `x` is less than the previous element, the stack is not sorted.
		if x < prev {
			return false
		}
		prev = x
	}

	// If the loop completes without finding any unsorted elements, the stack is sorted.
	return true
}
