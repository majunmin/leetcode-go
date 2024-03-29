package exercise_2024

// https://leetcode.cn/problems/sort-colors/?envType=study-plan-v2&envId=top-100-liked
func sortColors(nums []int) {
	l, r := 0, len(nums)-1
	var i int
	for i <= r {
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		}
		i++
	}
}
