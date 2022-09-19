package offer_042

//https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/?envType=study-plan&id=lcof
func maxSubArray(nums []int) int {
	return solution2(nums)
}

func solution2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	pre := nums[0]
	max := nums[0]
	for i := 1; i < n; i++ {
		if pre <= 0 {
			pre = nums[i]
		} else {
			pre += nums[i]
		}
		if pre > max {
			max = pre
		}
	}
	return max
}

func solution(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		if nums[i-1] <= 0 {
			nums[i] = nums[i]
		} else {
			nums[i] += nums[i-1]
		}
	}
	max := nums[0]
	for i := 0; i < n; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}
