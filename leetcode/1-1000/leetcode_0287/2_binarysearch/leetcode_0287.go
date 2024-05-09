package __binarysearch

// https://leetcode.cn/problems/find-the-duplicate-number/
// 题解:https://leetcode.cn/problems/find-the-duplicate-number/solutions/7038/er-fen-fa-si-lu-ji-dai-ma-python-by-liweiwei1419/
func findDuplicate(nums []int) int {
	var (
		n           = len(nums)
		left, right = 0, n // 根据题目的数据范围(开区间)
	)
	for left+1 < right {
		mid := left + (right-left)>>1
		// 统计nums[mid] 的个数
		var cnt int
		for _, num := range nums {
			if num <= nums[mid] {
				cnt++
			}
		}
		// 红蓝染色法:
		if cnt > mid {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}
