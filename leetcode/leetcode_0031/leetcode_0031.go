package leetcode_0031

// https://leetcode.cn/problems/next-permutation/
func nextPermutation(nums []int) {
	// param check
	n := len(nums)
	if n <= 1 {
		return
	}

	left := n - 2
	for ; left >= 0; left-- {
		if nums[left] < nums[left+1] {
			break
		}
	}
	// 说明当前  数组排列是最大值,直接返回
	if left == -1 {
		reverse(&nums, 0, n-1)
		return
	}

	// 1. 找到  右边第一个 大于 nums[left] 的index right
	right := n - 1
	for ; right > left; right-- {
		if nums[right] > nums[left] {
			// swap
			break
		}
	}
	// 2. swap left | right
	nums[left], nums[right] = nums[right], nums[left]

	//  3. reverse  [left + 1, n)
	reverse(&nums, left+1, n-1)
}

func reverse(nums *[]int, left, right int) {
	for left < right {
		(*nums)[left], (*nums)[right] = (*nums)[right], (*nums)[left]
		left++
		right--
	}
}
