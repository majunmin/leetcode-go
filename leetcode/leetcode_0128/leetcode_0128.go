package leetcode_0128

// https://leetcode.cn/problems/longest-consecutive-sequence/?envType=study-plan-v2&envId=top-100-liked
func longestConsecutive(nums []int) int {
	return dpSolution(nums)
}

func dpSolution(nums []int) int {
	boundMap := make(map[int]int)
	var result int
	for _, num := range nums {
		if _, exist := boundMap[num]; exist {
			continue
		}
		// default is zero
		left, right := boundMap[num-1], boundMap[num+1]
		curLen := left + right + 1
		// update result
		result = max(result, curLen)
		// 将num加入map中，表示已经遍历过该值。其对应的value可以为任意值。
		boundMap[num] = -1
		// 更新当前连续区间左边界和右边界对应的区间长度
		boundMap[num-left] = curLen
		boundMap[num+right] = curLen
	}

	return result
}

// 使用 map 记录右边界，同时更新
func solution1(nums []int) int {
	// 记录每个 num 可以达到最大的右边界
	rightBound := make(map[int]int)
	for _, n := range nums {
		rightBound[n] = n
	}
	var result int
	// 更新右边界
	for _, num := range nums {
		right := rightBound[num]
		if _, exist := rightBound[num-1]; exist {
			continue
		}
		_, exist := rightBound[right+1]
		for exist {
			right++
			_, exist = rightBound[right+1]
		}
		rightBound[num] = right
		result = max(result, right-num+1)
	}
	return result
}
