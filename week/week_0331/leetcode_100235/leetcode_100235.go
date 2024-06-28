package leetcode_100235

func maxBottlesDrunk(numBottles int, numExchange int) int {
	// param check
	if numBottles == 0 || numExchange > numBottles {
		return numBottles
	}
	var (
		total        = numBottles
		emptyBottles = numBottles
	)
	for emptyBottles >= numExchange {
		// exchange
		emptyBottles = emptyBottles - numExchange
		numExchange++
		total += 1
		emptyBottles += 1
	}
	return total
}
