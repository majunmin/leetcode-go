package leetcode_00291

func numberOfSpecialChars(word string) int {
	var (
		lowerIdx = make([]int, 26)
		upperIDx = make([]int, 26)
		n        = len(word)
		result   int
	)
	for i := 0; i < 26; i++ {
		lowerIdx[i] = -1
		lowerIdx[i] = -1
	}
	for i := range word {
		if word[i] >= 'a' && word[i] <= 'z' {
			lowerIdx[word[i]-'a'] = i
		} else {
			upperIDx[word[n-1-i]-'A'] = i
		}
	}
	for i := 0; i < 26; i++ {
		if upperIDx[i] == -1 || lowerIdx[i] == -1 {
			continue
		}
		if upperIDx[i] > lowerIdx[i] {
			result++
		}
	}
	return result
}
