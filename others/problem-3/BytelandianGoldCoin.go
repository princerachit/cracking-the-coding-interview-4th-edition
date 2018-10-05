package problems

var dp map[uint]uint64

//ByteLandianGoldCoinProblem solves https://www.spoj.com/submit/COINS/
func ByteLandianGoldCoinProblem(n int64) uint64 {
	if n <= 0 {
		return 0
	}

	//Create map with key of uint and valueType of uint64
	dp = make(map[uint]uint64)
	return exchangeCoin(uint(n))
}

func exchangeCoin(n uint) uint64 {

	if n <= 0 {
		return 0
	}

	// return exchange of n, if we already have calculated the exchange value against n.
	if dp[n] != 0 {
		return dp[n]
	}

	// Exchange n into n/2, n/3, and n/4. exchangeCoin will further divide the passed value
	// e.g:  12 into 6, 4 and 3. Exchange coin will see if it can get sum greater than 6
	// by diving 6 into 3, 2, and 0 which is 5 so it will return 6 instead.
	var aux = exchangeCoin(n/2) + exchangeCoin(n/3) + exchangeCoin(n/4)

	// If the sum of coins is greater than the n then keep the sum else keep n in the map
	// so that we can reuse it.
	if aux > uint64(n) {
		dp[n] = aux
	} else {
		dp[n] = uint64(n)
	}

	return dp[n]
}
