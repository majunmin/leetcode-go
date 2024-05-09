package leetcode_2105

// https://leetcode.cn/problems/watering-plants-ii/
// 哪里有问题？
func minimumRefill(plants []int, capacityA int, capacityB int) int {
	var (
		result           int
		left, right      = 0, len(plants) - 1
		lremain, rremain = capacityA, capacityB
	)
	for left < right {
		for lremain >= plants[left] {
			lremain -= plants[left]
			left++
		}
		for rremain >= plants[right] {
			rremain -= plants[right]
			right--
		}
		if left < right {
			result += 2
			lremain, rremain = capacityA, capacityB
		} else if left == right {
			result += 1
		}
	}
	return result
}

func minimumRefill2(plants []int, capacityA int, capacityB int) int {
	var (
		result           int
		left, right      = 0, len(plants) - 1
		lremain, rremain = capacityA, capacityB
	)
	for left <= right {
		if left == right {
			// 需要在词灌一下水
			if lremain < plants[left] && rremain < plants[right] {
				result++
			}
			break
		}
		if lremain < plants[left] {
			lremain = capacityA
			result++
		}
		lremain -= plants[left]
		// 移动指针
		left++

		if rremain < plants[right] {
			rremain = capacityB
			result++
		}
		rremain -= plants[right]
		// 移动指针
		right--
	}
	return result
}
