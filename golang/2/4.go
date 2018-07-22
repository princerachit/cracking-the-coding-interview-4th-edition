package problems

// You have two numbers represented by a linked list, where each node contains a sin-
// gle digit. The digits are stored in reverse order, such that the 1’s digit is at the head of
// the list. Write a function that adds the two numbers and returns the sum as a linked
// list.
// EXAMPLE
// Input: (3 -> 1 -> 5) + (5 -> 9 -> 2)
// Output: 8 -> 0 -> 8

type node struct {
	data int
	next *node
}

func calcSumFromLinkedList(l1, l2 *node) *node {
	// iterators to iterate over elements of both lists
	iter1 := l1
	iter2 := l2
	// sumList contains sum of the given two lists
	sumList := &node{next: nil, data: -1}
	// iterator over sumList
	curSumList := sumList

	// carry will hold the carry of sums
	var carry int

	// keep performing sum until there exist an unprocessed
	// element left in either of the list
	// or there is a carry from previous addition
	for iter1 != nil || iter2 != nil || carry != 0 {
		digSum := 0
		// add cur digit and advance the pointer
		if iter1 != nil {
			digSum += iter1.data
			iter1 = iter1.next
		}
		// add cur digit and advance the pointer
		if iter2 != nil {
			digSum += iter2.data
			iter2 = iter2.next
		}
		// add carry from previous sum
		digSum += carry
		// calculate new carry
		carry = digSum / 10
		// update the sum digit
		curSumList.data = digSum % 10
		// create new node for next iteration
		curSumList.next = &node{}
		curSumList = curSumList.next
	}
	// drop the extra node at the end of list
	curSumList.next = nil
	return sumList
}
