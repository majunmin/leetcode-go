package leetcode_0167

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/?envType=study-plan-v2&envId=top-interview-150
func twoSum(numbers []int, target int) []int {
	// param check
	if len(numbers) < 2 {
		return nil
	}
	left, right := 0, len(numbers)-1
	for left < right {
		if numbers[left]+numbers[right] < target {
			left++
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			return []int{left + 1, right + 1}
		}
	}
	// left == right
	return nil
}
