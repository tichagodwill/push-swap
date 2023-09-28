package funcs

import (
	"sort"
)

// SmallSort takes a slice of integers 'nums', sorts it, and returns a string of instructions to achieve the sorting.
func Sort(nums []int) string {
	res := ""
	res = SmallInstructions(nums)
	return res
}

// SmallInstructions takes a slice of integers 'nums' and returns a string of instructions to sort it.
func SmallInstructions(nums []int) string {
	half := int(len(nums) / 2)
	res := ""

	// First loop: Sorting elements in StackA
	for {
		a, b := StackA, StackB
		if isASorted(a) {
			break
		}
		op := match(aStack(a, half), a, b)
		ApplyInsruction(op)
		res += op + "\n"
	}

	// Second loop: Moving elements from StackB to StackA
	for {
		a, b := StackA, StackB
		if len(a) == len(nums) && isASorted(a) {
			break
		}
		ins := bStack(b)
		ApplyInsruction(ins)
		res += ins + "\n"
	}
	return res
}

// aStack implements stack operations for StackA.
func aStack(a []int, half int) string {
	swap := saIdx(a)
	mid, i := (len(a) / 2), (len(a) - 1) // i is the top of the stack
	if a[i] < half {
		return "pb"
	}
	if swap > 0 {
		if inOrder(a[i-1], a[i]) { // a[i] > a[i-1] and the difference = 1
			return "sa"
		} else if swap < mid {
			return "rra"
		}
	} else {
		minNumIdx := smallest(a)
		if minNumIdx == i { // the minimum is in the top of the stack // will work only if the stack has two elements
			return "pb"
		} else if minNumIdx < mid {
			return "rra"
		} else if inOrder(a[i], a[0]) { // a[0] > a[i] and the difference = 1
			return "rra"
		}
	}
	return "ra"
}

// bStack implements stack operations for StackB.
func bStack(bstack []int) string {
	b := make([]int, len(bstack))
	copy(b, bstack)
	swap := sbIdx(b)
	mid, j := (len(b) / 2), (len(b) - 1)
	if swap > 0 {
		if inOrder(b[j], b[j-1]) {
			return "sb"
		} else if swap < mid {
			return "rrb"
		}
	} else {
		maxNumIdx := biggest(b)
		if maxNumIdx == j {
			return "pa"
		} else if maxNumIdx < mid || inOrder(b[0], b[j]) {
			return "rrb"
		}
	}
	return "rb"
}

// match determines the appropriate instruction to apply based on conditions.
func match(ins string, a, b []int) string {
	if len(b) < 2 || isBSorted(b) {
		return ins
	}
	bins := bStack(b)
	switch ins { // ins is an instruction from aStack(a)
	case "pb":
		if bins == "sb" {
			if t := len(a) - saIdx(a); t < 3 {
				if t == 1 {
					return "ss"
				}
				return "ra"
			} else {
				return "sb"
			}
		}
	case "sa":
		if t := len(b) - sbIdx(b); t < 3 {
			if bins == "sb" {
				return "ss"
			}
			return "rb"
		}
	case "ra":
		if bins == "rb" || bins == "sb" && len(b) == 2 {
			return "rr"
		}
	case "rra":
		if bins == "rrb" || bins == "sb" && len(b) == 2 {
			return "rrr"
		}
	}
	return ins
}

// smallest finds the index of the smallest element in 'nums'.
func smallest(nums []int) int {
	min, idx := 0x3f3f3f3f, -1
	for i := 0; i < len(nums); i++ {
		if min > int(nums[i]) {
			min = int(nums[i])
			idx = i
		}
	}
	return idx
}

// biggest finds the index of the largest element in 'nums'.
func biggest(nums []int) int {
	max, idx := -1, -1
	for i := 0; i < len(nums); i++ {
		if max < int(nums[i]) {
			max = int(nums[i])
			idx = i
		}
	}
	return idx
}

// saIdx finds the index for an ascending sequence in 'nums'.
func saIdx(nums []int) int {
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i]-nums[i-1] == 1 {
			return i
		}
	}
	return -1
}

// sbIdx finds the index for a descending sequence in 'nums'.
func sbIdx(nums []int) int {
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i-1]-nums[i] == 1 {
			return i
		}
	}
	return -1
}

// inOrder returns true if 'a' and 'b' are in ascending order.
func inOrder(a, b int) bool {
	return b-a == 1
}

// isASorted checks if 'nums' is sorted in ascending order.
func isASorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]-nums[i+1] != 1 {
			return false
		}
	}
	return true
}

// isBSorted checks if 'nums' is sorted in descending order.
func isBSorted(nums []int) bool {
	// Create an empty slice of int
	intSlice := make([]int, len(nums))

	// Perform the conversion element-wise
	for i, v := range nums {
		intSlice[i] = int(v)
	}

	return sort.IntsAreSorted(intSlice)
}

// UInts sorts 'nums' and updates it with the sorted values.
func UInts(nums []int) {
	copiedNums := make([]int, len(nums))
	copy(copiedNums, nums)
	sort.Ints(copiedNums)
	for i, n := range nums {
		nums[i] = IndexSearch(copiedNums, n)
	}
}

// IndexSearch performs a search on 'nums' and returns the index of 'target'.
func IndexSearch(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		}
	}
	return -1
}
