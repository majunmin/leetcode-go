package leetcode_2588

// https://leetcode.cn/problems/count-the-number-of-beautiful-subarrays/
func beautifulSubarrays(nums []int) int64 {
	var cnt int
	for i := 0; i < len(nums); i++ {
		var (
			curCnt int
			cur    = nums[i]
		)
		if cur == 0 {
			curCnt++
		}
		for j := i + 1; j < len(nums); j++ {
			cur ^= nums[j]
			if cur == 0 {
				curCnt++
			}
		}
		cnt += curCnt
	}
	return int64(cnt)
}

// 使用前缀和 避免重复计算
