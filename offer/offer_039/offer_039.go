package offer_039

// https://leetcode.cn/problems/shu-zu-zhong-chu-xian-ci-shu-chao-guo-yi-ban-de-shu-zi-lcof/?envType=study-plan&id=lcof
func majorityElement(nums []int) int {
	// 分治算法
	// param check
	if len(nums) == 0 {
		return -1
	}
	return majority(nums, 0, len(nums)-1)
}

func majority(nums []int, left int, right int) int {
	// terminate
	if left >= right {
		return nums[left]
	}

	// 分而治之
	mid := left + (right-left)>>1
	leftVal := majority(nums, left, mid)
	rightVal := majority(nums, mid+1, right)

	// 归并结果
	if leftVal == rightVal {
		return leftVal
	}
	return findMajority(nums, left, right, leftVal, rightVal)
}

func findMajority(nums []int, left int, right int, leftVal int, rightVal int) int {
	var count int
	for i := left; i <= right; i++ {
		if nums[i] == leftVal {
			count++
		}
	}
	if count > (right-left+1)>>1 {
		return leftVal
	}
	return rightVal
}

// Boyer-Moore 投票算法
func solution1(nums []int) int {

	// param check
	if len(nums) == 0 {
		return -1
	}

	// init state
	candidate, count := nums[0], 0
	for _, num := range nums {
		if num == candidate {
			count += 1
			continue
		}

		count--
		if count == 0 {
			candidate = num
		}
	}
	return candidate
}
