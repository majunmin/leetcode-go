package offer_003

//https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/?envType=study-plan&id=lcof
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == i {
			continue
		}

		//出现重复的条件.  i != nums[i] && nums[i] == nums[nums[i]]
		if nums[i] == nums[nums[i]] {
			return nums[i]
		}
		// swap nums[i], nums[nums[i]]
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		i--
	}
	return -1
}

func mapSolution(nums []int) int {
	m := make(map[int]struct{})
	for n := range nums {
		if _, exist := m[n]; exist {
			return n
		}
		m[n] = struct{}{}
	}
	return -1
}
