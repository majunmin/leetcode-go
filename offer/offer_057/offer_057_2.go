package offer_057

// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/?envType=study-plan&id=lcof
func findContinuousSequence(target int) [][]int {
	// param check
	if target < 1 {
		return nil
	}
	result := make([][]int, 0)
	// 滑动窗口解法
	left, right := 1, 1 // 从 1 开始 正整数
	sum := 1
	for left < right || left == 1 {
		if sum == target {
			// add to result
			res := make([]int, 0, right-left+1)
			for i := left; i <= right; i++ {
				res = append(res, i)
			}
			result = append(result, res)
			right++
			sum += right
		} else if sum < target {
			right++
			sum += right
		} else {
			sum -= left
			left++
		}
	}
	return result
}
