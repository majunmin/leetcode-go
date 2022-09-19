package offer_050

//https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/?envType=study-plan&id=lcof
func firstUniqChar(s string) byte {
	n := len(s)
	// 保存 字符和索引的映射关系(重复字符索引为 -1)
	m := make(map[byte]int)
	for i, b := range []byte(s) {
		// 重复的 字符  替换为 -1
		if _, exist := m[b]; exist {
			m[b] = -1
			continue
		}
		m[b] = i
	}
	first := n
	for _, idx := range m {
		if idx != -1 && idx < first {
			first = idx
		}
	}
	if first == n {
		return ' '
	}
	return s[first]
}

func mapSolution(s string) byte {
	m := make(map[byte]int)
	for _, b := range []byte(s) {
		m[b]++
	}
	for b, cnt := range m {
		if cnt == 1 {
			return b
		}
	}
	return ' '
}
