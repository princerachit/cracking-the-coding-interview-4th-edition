package problems

import "fmt"

/* Imagine a (literal) stack of plates. If the stack gets too high, it might topple. There-
fore, in real life, we would likely start a new stack when the previous stack exceeds
some threshold. Implement a data structure SetOfStacks that mimics this. SetOf-
Stacks should be composed of several stacks, and should create a new stack once
the previous one exceeds capacity. SetOfStacks.push() and SetOfStacks.pop() should
behave identically to a single stack (that is, pop() should return the same values as it
would if there were just a single stack).
FOLLOW UP
Implement a function popAt(int index) which performs a pop operation on a specific
sub-stack */

// Internal stack
type Stack struct {
	array []int
	top   int
}

// SetOfStacks which internally utilizes Stack
type SetOfStacks struct {
	activeStack int
	stacks      []Stack
}

// NewSetOfStacks is a constructor for SetOfStacks
func NewSetOfStacks() *SetOfStacks {
	return &SetOfStacks{
		activeStack: -1,
	}
}

// NewStack is a constructor for Stack
func NewStack() *Stack {
	return &Stack{array: make([]int, 10), top: -1}
}

func (ss *SetOfStacks) push(val int) {
	// if no stack is present then create one
	if ss.stacks == nil || len(ss.stacks) == 0 {
		ss.stacks = append(ss.stacks, *NewStack())
		ss.activeStack = 0
	}

	// check the length of the stacks
	// if the length of last Stack is 10 then
	// create a new Stack
	// pushing element to the last stack

	if ss.stacks[ss.activeStack].top == 9 {
		if len(ss.stacks)-1 == ss.activeStack {
			ss.stacks = append(ss.stacks, *NewStack())
		}
		ss.activeStack++
	}

	lastStack := ss.stacks[ss.activeStack]

	lastStack.array[lastStack.top+1] = val
	// increment top by 1
	lastStack.top = lastStack.top + 1
}

func (ss *SetOfStacks) pop() (int, error) {

	// check length and activeStack
	if ss.activeStack == -1 {
		return 0, fmt.Errorf("stack underflow")
	}

	// find the last stack
	lastStack := ss.stacks[ss.activeStack]
	// retrieve top val from the last stack
	poppedVal := lastStack.array[lastStack.top]
	lastStack.top--
	if lastStack.top == -1 {
		ss.activeStack--
	}
	return poppedVal, nil
}
