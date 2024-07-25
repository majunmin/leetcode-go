package leetcode_3096

// https://leetcode.cn/problems/minimum-levels-to-gain-more-points/?envType=daily-question&envId=2024-07-19
func minimumLevels(possible []int) int {
	var (
		pre   = make([]int, len(possible)+1)
		total int // 总分,
	)
	// 如果 Alice 拿到 a 分, Bob 的得分 = total - a
	for i, p := range possible {
		if p == 0 {
			total -= 1
		} else {
			total += 1
		}
		pre[i+1] = total
	}
	for i := 1; i < len(pre); i++ {
		if pre[i] > total-pre[i] {
			return i
		}
	}
	return -1
}
