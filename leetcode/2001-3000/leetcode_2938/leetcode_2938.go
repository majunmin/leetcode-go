package leetcode_2938

// https://leetcode.cn/problems/separate-black-and-white-balls/
func minimumSteps(s string) int64 {
	var (
		cntOne int
		result int
	)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			cntOne++
			continue
		}
		// 遇到0统计 移动次数
		result += cntOne
	}
	return int64(result)
}
