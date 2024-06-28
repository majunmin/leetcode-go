package leetcode_0520

// https://leetcode.cn/problems/detect-capital/?envType=daily-question&envId=2024-06-23
func detectCapitalUse(word string) bool {
	if isUpper(word[0]) {
		return same(word[1:])
	}
	return same(word)
}

func same(word string) bool {
	b := isUpper(word[0])
	for i := 1; i < len(word); i++ {
		if b != isUpper(word[i]) {
			return false
		}
	}
	return true
}

func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}
