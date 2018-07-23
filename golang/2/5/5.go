package problems

// Given a circular linked list, implement an algorithm which returns node at the begin-
// ning of the loop.
// DEFINITION
// Circular linked list: A (corrupt) linked list in which a node’s next pointer points to an
// earlier node, so as to make a loop in the linked list.
// EXAMPLE
// input: A -> B -> C -> D -> E -> C [the same C as earlier]
// output: C

type linkedList struct {
	data rune
	next *linkedList
}

// function does not check for the non looping linkedlist
// improve later
func findLoop(list *linkedList) rune {
	// slow iter moves one step
	slowIter := list
	// fast iter moves two steps
	fastIter := list

	// loop breaks only when both iters meet
	for {
		slowIter = slowIter.next
		fastIter = fastIter.next.next
		if slowIter == fastIter {
			break
		}
	}

	// reset slow iter to head
	slowIter = list
	// the two iters will meet at the start point of the loop
	for slowIter != fastIter {
		// move both iters at same speed
		slowIter = slowIter.next
		fastIter = fastIter.next
	}

	return slowIter.data
}
