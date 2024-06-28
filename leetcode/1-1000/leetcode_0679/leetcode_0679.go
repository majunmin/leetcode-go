package leetcode_0679

// https://leetcode.cn/problems/24-game/
func judgePoint24(cards []int) bool {
	var (
		nums = make([]float64, 0, len(cards))
	)

	for _, card := range cards {
		nums = append(nums, float64(card))
	}

	allPermutes := permute(nums)
	for _, permutes := range allPermutes {
		res := dfs(permutes, 0, 3)
		for item := range res {
			if abs(item-24.0) <= 1e-7 {
				return true
			}
		}
	}
	return false
}

func abs(num float64) float64 {
	if num < 0 {
		return -num
	}
	return num
}

// nums[start:end+1] + 不同的操作符可以组成的数字
// 统计可以组成的数字
func dfs(nums []float64, start, end int) map[float64]bool {
	var result = make(map[float64]bool)
	// termiante
	if start == end {
		result[nums[start]] = true
		return result
	}
	// 切分数字
	for i := start; i < end; i++ {
		res1 := dfs(nums, start, i)
		res2 := dfs(nums, i+1, end)
		for num1 := range res1 {
			for num2 := range res2 {
				result[num1+num2] = true
				result[num1-num2] = true
				result[num1*num2] = true
				result[num1/num2] = true
			}
		}
	}
	return result
}

func permute(nums []float64) [][]float64 {
	var (
		result  [][]float64
		visited = make([]bool, len(nums))
		dfs     func([]float64)
	)
	dfs = func(path []float64) {
		// terminate
		if len(path) == len(nums) {
			temp := make([]float64, len(path))
			copy(temp, path)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			path = append(path, nums[i])
			dfs(path)
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
	dfs([]float64{})
	return result
}
