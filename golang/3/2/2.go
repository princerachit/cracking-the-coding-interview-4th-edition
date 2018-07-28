package problems

import "fmt"

// How would you design a stack which, in addition to push and pop, also has a function
// min which returns the minimum element? Push, pop and min should all operate in
// O(1) time.
type element struct {
	data          int
	smallestAfter int
}

type stack []element

// pop will return the pop an element from the stack and return it along with the new
// top position
func pop(s stack, top int) (*element, int, error) {
	if top == -1 {
		return nil, -1, fmt.Errorf("underflow")
	}
	poppedElement := s[top]
	top--
	return &poppedElement, top, nil
}

// push will push the new element on the stack and return the new
// top value
func push(s stack, e element, top int) int {
	top++
	s[top] = e
	if top == 0 {
		s[top].smallestAfter = e.data
	} else {
		if e.data < s[top-1].data {
			s[top].smallestAfter = e.data
		} else {
			s[top].smallestAfter = s[top-1].data
		}
	}
	return top
}

// minElement will return the minimum element on the stack
func minElement(s stack, top int) (int, error) {
	if top < 0 {
		return -1, fmt.Errorf("empty stack")
	}
	return s[top].smallestAfter, nil
}
