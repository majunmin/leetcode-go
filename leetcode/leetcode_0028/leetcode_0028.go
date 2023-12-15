package leetcode_0028

// https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
func strStr(haystack string, needle string) int {
	return kmpSolution(haystack, needle)
}

func kmpSolution(haystack string, needle string) int {
	next := buildNext(needle)
	i, j := 0, 0
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else if j > 0 {
			j = next[j-1]
		} else {
			// 从头开始比， j = 0
			i++
		}
		if j == len(needle) {
			return i - j
		}
	}
	return -1
}

func buildNext(needle string) []int {
	var next []int
	next = append(next, 0)
	i, prefixLen := 1, 0
	for i < len(needle) {
		if needle[prefixLen] == needle[i] {
			prefixLen += 1
			next = append(next, prefixLen)
			i += 1
		} else {
			if prefixLen == 0 {
				next = append(next, 0)
				i++
			} else {
				prefixLen = next[prefixLen-1]
			}
		}
	}
	return next
}

func boyerMooreSearch(text, pattern string) int {
	n := len(text)    // 文本串长度
	m := len(pattern) // 模式串长度

	if m == 0 {
		return 0
	}

	// 构建坏字符规则表
	badCharTable := make(map[byte]int)
	for i := 0; i < m-1; i++ {
		// 计算模式串中每个字符到末尾字符的距离，作为坏字符规则表的值
		badCharTable[pattern[i]] = m - 1 - i
	}

	// 构建好后缀规则表
	suffixTable := make([]int, m)
	prefix := make([]bool, m)
	for i := 0; i < m; i++ {
		suffixTable[i] = -1
		prefix[i] = false
	}
	for i := 0; i < m-1; i++ {
		j := i
		k := 0
		for j >= 0 && pattern[j] == pattern[m-1-k] {
			// 从模式串末尾往前匹配，记录每个长度的好后缀的起始位置
			j--
			k++
			suffixTable[k] = j + 1
		}
		if j == -1 {
			// 如果匹配到模式串的起始位置，说明当前长度的好后缀是一个前缀
			prefix[k] = true
		}
	}

	// 开始搜索
	i := m - 1 // 文本串中当前比对的位置
	for i < n {
		j := m - 1 // 模式串中当前比对的位置

		// 从后往前比较字符，直到模式串结束或者出现不匹配的字符
		for j >= 0 && text[i] == pattern[j] {
			i--
			j--
		}

		if j == -1 {
			// 如果模式串比对完毕，说明匹配成功，返回起始位置
			return i + 1
		}

		// 根据坏字符规则和好后缀规则计算移动的位数
		badCharShift := badCharTable[text[i]] // 坏字符规则表中当前字符的移动位数
		goodSuffixShift := 0
		if j < m-1 {
			// 如果存在好后缀，则根据好后缀规则计算移动位数
			goodSuffixShift = moveByGoodSuffix(j, m, suffixTable, prefix)
		}

		// 取坏字符规则和好后缀规则中较大的移动位数
		shift := max(badCharShift, goodSuffixShift)
		i += shift
	}

	return -1 // 未找到匹配
}

// moveByGoodSuffix 根据好后缀规则计算移动的位数。
func moveByGoodSuffix(j, m int, suffixTable []int, prefix []bool) int {
	k := m - 1 - j // 好后缀长度
	if suffixTable[k] != -1 {
		// 如果好后缀在好后缀规则表中存在，直接返回移动位数
		return j - suffixTable[k] + 1
	}
	for r := j + 2; r <= m-1; r++ {
		if prefix[m-r] {
			// 如果存在与好后缀匹配的前缀，则返回移动位数
			return r
		}
	}
	// 如果没有匹配的前缀，则整个好后缀都不匹配，返回模式串长度作为移动位数
	return m
}
