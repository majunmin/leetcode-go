package leetcode_0503

// https://leetcode.cn/problems/next-greater-element-ii/?envType=daily-question&envId=2024-06-24
func nextGreaterElements(nums []int) []int {
	// 使用 一个单调递减栈, 从末尾向前遍历。
	// 最后一个元素单独处理.(循环向后找)
	if len(nums) == 1 {
		return []int{-1}
	}
	var (
		n      = len(nums)
		result = make([]int, n)
		stack  = make([]int, 0, n)
	)
	stack = append(stack, nums[n-1])
	for i := n - 2; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1]
		} else {
			result[i] = -1
		}
		stack = append(stack, nums[i])
	}
	for i := n - 1; i > 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			result[i] = stack[len(stack)-1]
		} else {
			result[i] = -1
		}
		stack = append(stack, nums[i])
	}
	return result
}
