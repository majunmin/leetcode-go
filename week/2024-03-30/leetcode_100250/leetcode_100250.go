package leetcode_100250

// https://leetcode.cn/contest/biweekly-contest-127/problems/minimum-levels-to-gain-more-points/
func minimumLevels(possible []int) int {
	// param check
	if len(possible) == 0 {
		return -1
	}
	var (
		total int
		step  int
		score int
	)
	for i := 0; i < len(possible); i++ {
		total += getScore(possible[i])
	}
	for i := 0; i < len(possible)-1; i++ {
		step++
		score += getScore(possible[i])
		total -= getScore(possible[i])
		if score > total {
			return step
		}
	}
	return -1
}

func getScore(i int) int {
	if i == 0 {
		return -1
	}
	return 1
}
