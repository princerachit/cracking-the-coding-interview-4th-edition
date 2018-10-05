package problem2

//knapSack - Returns the maximum value that can be put in a knapsack of capacity w;
//val and wt which represent values and weights associated with items respectively
func knapSack(w int, wt []int, val []int) int {
	//recursion breaker
	//knapSack is full or there is no new items
	if w <= 0 || len(wt) == 0 || len(val) == 0 {
		return 0
	}
	//if item is to heavy - pass
	if wt[0] > w {
		return knapSack(w, wt[1:], val[1:])
	}

	// return the maximum of two cases:
	// 1st item included
	// not included
	addFirstItem := val[0] + knapSack(w-wt[0], wt[1:], val[1:])
	dropFirstItem := knapSack(w, wt[1:], val[1:])

	if addFirstItem > dropFirstItem {
		return addFirstItem
	}
	return dropFirstItem
}
