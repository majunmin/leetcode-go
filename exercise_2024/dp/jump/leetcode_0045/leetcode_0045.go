package leetcode_0045

func jump(nums []int) int {
	return solution2(nums)
}

func solution2(nums []int) int {
	var maxPosition int
	var steps int
	var end int
	for i := 0; i < len(nums); i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

// 反向跳 jump
func solution1(nums []int) int {
	// param check
	if len(nums) == 0 {
		panic("invalid param")
	}
	if len(nums) == 1 {
		return 0
	}

	var steps int
	position := len(nums) - 1
	for position > 0 {
		for i := 0; i < position; i++ {
			if i+nums[i] > position {
				steps++
				position = i
				break
			}
		}
	}
	return steps
}
