package offer_061

import "sort"

//https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/?envType=study-plan&id=lcof
//是顺子的条件是:
// 1.  抽到 的 牌  除了0 之外 没有重复
// 2. 除了0 之外 剩余的牌的 最大值 与最小值的差 <5
func isStraight(nums []int) bool {

	return sortSolution(nums)
}

// 排序后 从第一个不为0的数字 看 , 后面的数据无重复 && max - min <5
func sortSolution(nums []int) bool {

	// param check

	length := len(nums)
	if length != 5 {
		return false
	}

	// sort
	sort.Ints(nums)

	min, max := nums[length-1], nums[length-1]
	for i := length - 2; i >= 0; i-- {
		if nums[i] == 0 {
			break
		}
		if nums[i] == nums[i+1] { // 如果存在重复值
			return false
		}

		min = nums[i]
	}

	return max-min < 5

}

// 用一个 set 结构用于去重
func setSolution(nums []int) bool {
	if len(nums) != 5 {
		return false
	}

	m := make(map[int]struct{})
	min, max := 14, 0
	for _, num := range nums {
		if num == 0 {
			continue
		}

		// 如果 存在 重复的牌,返回 false
		if _, exist := m[num]; exist {
			return false
		}

		//  update min & max
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return max-min < 5
}
