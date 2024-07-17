package lcr_005

// https://leetcode.cn/problems/aseY1I/
func maxProduct(words []string) int {
	var (
		//words[i] 仅包含小写字母，26位,可以利用位运算快速判断 两个单词是否包含重复字符
		masks  = make([]int, 0, len(words))
		result int
	)
	for _, word := range words {
		var m int
		for i := 0; i < len(word); i++ {
			m |= 1 << (word[i] - 'a')
		}
		masks = append(masks, m)
	}
	//
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			// 无重复字符
			if masks[i]&masks[1] == 0 {
				result = max(result, len(words[i])*len(words[j]))
			}
		}
	}
	return result
}
