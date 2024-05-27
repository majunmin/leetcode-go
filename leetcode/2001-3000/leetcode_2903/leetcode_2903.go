package leetcode_2903

// https://leetcode.cn/problems/find-indices-with-index-and-value-difference-i/?envType=daily-question&envId=2024-05-27
func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	// 暴力法
	n := len(nums)
	for i := 0; i < n-indexDifference; i++ {
		for j := indexDifference; j < n; j++ {
			if abs(nums[i]-nums[j]) >= valueDifference {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func abs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}
