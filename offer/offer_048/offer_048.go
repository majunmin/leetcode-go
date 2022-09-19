package offer_048

//https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/?plan=lcof&plan_progress=z2ptdrc
func lengthOfLongestSubstring(s string) int {
	return slidingWindowSolution(s)
}

//
func slidingWindowSolution(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return n
	}
	maxLen := 0
	left, right := 0, 0
	m := make(map[byte]struct{})
	for right < n {
		if _, exist := m[s[right]]; !exist {
			maxLen = maxInt(maxLen, right-left+1)
			m[s[right]] = struct{}{}
			right++
			continue
		}

		// 如果 发现重复,就 缩小窗口
		if left < right {
			delete(m, s[left])
			left++
		}
	}
	return maxLen
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
