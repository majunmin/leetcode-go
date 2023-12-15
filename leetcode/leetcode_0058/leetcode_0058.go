package leetcode_0058

func lengthOfLastWord(s string) int {
	var result int
	var findWord bool
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if findWord {
				break
			}
			continue
		}

		findWord = true
		result++
	}
	return result
}
