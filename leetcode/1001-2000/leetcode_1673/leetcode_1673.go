package leetcode_1673

func mostCompetitive(nums []int, k int) []int {
	// param check
	if len(nums) < k {
		panic("invlaid param")
	}
	// 单调非递减栈
	var (
		n      = len(nums)
		stack  = make([]int, 0)
		result = make([]int, 0, k)
	)
	for i, num := range nums {
		//len(stack) > 0 && 栈上元素 + nums剩余元素 > k && curNum < 栈顶元素
		for len(stack) > 0 && len(stack)+n-i > k && num < stack[len(stack)-1] {
			// pop tail
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, num)
	}
	return result[:k]
}
