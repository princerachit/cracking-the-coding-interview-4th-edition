package problems

func printExecutionSequence(number int64) string {
	if number == 1 || number == 0 {
		return ""
	} else {
		if number%2 == 1 || number%2 == -1 {
			return "L" + printExecutionSequence((number+1)/2)
		} else {
			return "R" + printExecutionSequence(number/2)
		}
	}
}
