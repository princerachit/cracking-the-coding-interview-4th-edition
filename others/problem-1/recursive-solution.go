package problems

func printExecutionSequence(number int64) string {
	if number == 1 || number == 0 {
		return ""
	}
	if number%2 == 1 || number%2 == -1 {
		return "L" + printExecutionSequence((number+1)/2)
	}
	return "R" + printExecutionSequence(number/2)
}
