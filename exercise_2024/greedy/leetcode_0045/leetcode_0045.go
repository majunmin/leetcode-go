package leetcode_0045

// https://leetcode.cn/problems/jump-game-ii/description/?envType=study-plan-v2&envId=top-100-liked
func jump(nums []int) int {
	// 从尾部向前查找
	n := len(nums)
	if n == 1 {
		return 0
	}
	//
	var (
		end   = n - 1
		steps int
	)
	for end > 0 {
		for i := 0; i < end; i++ {
			if i+nums[i] >= end {
				end = i
				steps++
				break
			}
		}
	}
	return steps
}

func solution2(nums []int) int {
	// param check
	n := len(nums)
	if n == 1 {
		return 0
	}
	var (
		end   int
		reach int
		steps int
	)
	// 到达边界时 steps++
	//一个小技巧:遍历数组时，我们不访问最后一个元素,这是因为在访问最后一个元素之前,我们的边界一定大于等于最后一个位置,否则就无法跳到最后一个位置了.
	// 防止不必要的跳跃
	for i := 0; i < n-1; i++ {
		reach = max(reach, i+nums[i])
		if i == end {
			end = reach
			steps++
		}
	}
	return steps
}
