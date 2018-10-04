package problems

var dp map[uint]uint64

//ByteIndianGoldCoinProblem solves https://www.spoj.com/submit/COINS/
func ByteIndianGoldCoinProblem(n int64) uint64 {
	if n <= 0 {
		return 0
	}

	dp = make(map[uint]uint64)
	return solveByteIndianGoldCoinProblem(uint(n))
}

func solveByteIndianGoldCoinProblem(n uint) uint64 {
	if n <= 0 {
		return 0
	}

	if dp[n] != 0 {
		return dp[n]
	}

	var aux = solveByteIndianGoldCoinProblem(n/2) +
		solveByteIndianGoldCoinProblem(n/3) +
		solveByteIndianGoldCoinProblem(n/4)

	if aux > uint64(n) {
		dp[n] = aux
	} else {
		dp[n] = uint64(n)
	}

	return dp[n]
}
