package leetcode_0003

// 1. 使用map 记录窗口内 字符出现的次数, 如果当前窗口中每个字符的个数都是1，则继续扩展窗口
// 2. 如果 窗口中的某个字符的个数 > 1, 则缩容窗口.

// https://leetcode.cn/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {

	// set 结构, 记录字符时候否出现过
	occ := make(map[byte]bool)

	var r, result int
	// 枚举左指针
	for i := 0; i < len(s); i++ {
		if i > 0 {
			occ[s[i-1]] = false
		}

		//  滑动右指针
		for ; r < len(s); r++ {
			if occ[s[r]] {
				break
			}
			occ[s[r]] = true
		}
		result = max(result, r-i)
	}
	return result
}

// 滑动窗口解法
func solution1(s string) int {
	// param check
	if len(s) < 2 {
		return len(s)
	}
	cnt := make(map[byte]int, 0)
	var result int
	var l, r int
	var flag bool
	for r < len(s) {
		for r < len(s) {
			cnt[s[r]]++
			r++
			if cnt[s[r-1]] > 1 {
				flag = true
				break
			}
		}

		// 窗口中存在重复字符
		// 不满足条件, 窗口缩容
		for flag && l < r {
			cnt[s[l]]--
			l++
			if cnt[s[l-1]] == 1 {
				// 统计结果
				flag = false
			}
		}
		result = max(result, r-l)
	}
	return result
}
