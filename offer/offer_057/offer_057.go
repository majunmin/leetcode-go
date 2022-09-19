package offer_057

// https://leetcode.cn/problems/he-wei-sde-liang-ge-shu-zi-lcof/?envType=study-plan&id=lcof
func twoSum(nums []int, target int) []int {
	// param check
	n := len(nums)
	if n <= 1 {
		return nil
	}
	i, j := nums[0], nums[n-1]
	for i < j {
		v1, v2 := nums[i], nums[j]
		if v1+v2 == target {
			// find result
			return []int{v1, v2}
		} else if v1+v2 < target {
			i++
		} else {
			j--
		}
	}
	// 无解
	return nil
}
