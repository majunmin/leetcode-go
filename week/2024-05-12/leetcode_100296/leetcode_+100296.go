package leetcode_100296

func findPermutationDifference(s string, t string) int {
	var (
		sPos   = make(map[byte]int)
		tPos   = make(map[byte]int)
		result int
	)
	for i := 0; i < len(s); i++ {
		sPos[s[i]] = i
		tPos[t[i]] = i
	}
	for i := 0; i < len(s); i++ {
		result += abs(tPos[s[i]] - sPos[s[i]])
	}
	return result

}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a

}
