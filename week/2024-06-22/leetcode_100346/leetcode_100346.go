package leetcode_100346

// 按照 0 1 分组计数

func minOperations(nums []int) int {
	var (
		target = 0
		cnt    int
		i      int
	)
	for ; i < len(nums); i++ {
		if nums[i] == target {
			cnt++
			target = 1 - target
		}
	}

	return cnt

}
