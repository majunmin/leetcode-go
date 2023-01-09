package leetcode_0915

// https://leetcode.cn/problems/partition-array-into-disjoint-intervals/
func partitionDisjoint(nums []int) int {
	// param check
	if len(nums) == 0 {
		return -1
	}

	leftMax, max, idx := nums[0], nums[0], 0
	for i := 1; i < len(nums); i++ {
		max = maxInt(max, nums[i])
		if nums[i] < leftMax {
			leftMax = max
			idx = i
		}
	}
	return idx + 1
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//  5 0 3 8 6
//  90,47,69,10,43,92,31,73,61,97
