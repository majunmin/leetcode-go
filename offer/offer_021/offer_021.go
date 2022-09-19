package offer_021

import "sort"

//https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/?envType=study-plan&id=lcof
func exchange(nums []int) []int {
	// param check
	if len(nums) <= 1 {
		return nums
	}

	//
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]&1 == 1 && nums[j]&1 == 0 {
			// swap nums i, j
			nums[i], nums[j] = nums[j], nums[i]
		}
		// find left first even
		for i < j && nums[i]&1 == 0 {
			i += 1
		}
		// find right first odd
		for i < j && nums[j]&1 == 1 {
			j -= 1
		}
	}
	return nums
}

func goSolution(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]&1 == 1 && nums[j]&1 == 0
	})
	return nums
}
