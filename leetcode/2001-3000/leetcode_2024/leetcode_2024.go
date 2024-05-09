package leetcode_2024

// https://leetcode.cn/problems/maximize-the-confusion-of-an-exam/
func maxConsecutiveAnswers(answerKey string, k int) int {
	var (
		ansMap = make(map[byte]int)
		l, r   int
		result int
	)
	ansMap['T'] = 0
	ansMap['F'] = 0

	for ; r < len(answerKey); r++ {
		c := answerKey[r]
		ansMap[c]++
		if valid(ansMap, k) {
			result = max(result, r-l+1)
			continue
		}
		// shrink window
		for ; l < r && !valid(ansMap, k); l++ {
			ansMap[answerKey[l]]--
		}
	}
	return result
}

func valid(cnts map[byte]int, k int) bool {
	for _, cnt := range cnts {
		if cnt <= k {
			return true
		}
	}
	return false
}
