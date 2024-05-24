package leetcode_1542

// https://leetcode.cn/problems/find-longest-awesome-substring/
func longestAwesome(s string) int {
	var (
		pre    int
		pos    = make(map[int]int) // 记录前缀异或和出现的位置
		result int
	)
	pos[0] = -1
	for i := 0; i < len(s); i++ {
		pre ^= 1 << (s[i] - '0')
		for j := 0; j < 10; j++ {
			// 奇数
			if preIdx, exist := pos[pre^(1<<j)]; exist {
				result = max(result, i-preIdx)
			}
		}
		// 偶数
		if preIdx, exist := pos[pre]; exist {
			result = max(result, i-preIdx)
		}
		if _, exist := pos[pre]; !exist {
			pos[pre] = i
		}
	}
	return result
}

// 如何才算是是一个 回文字符串.
// - 所有字符出现次数都是偶数
// - 只有一个字符出现次数是奇数, 其他字符出现次数是偶数
// a ^ (*) = b.  ===> * = a ^ b
