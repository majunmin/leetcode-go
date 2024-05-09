package leetcode_2079

// https://leetcode.cn/problems/watering-plants/?envType=daily-question&envId=2024-05-08
func wateringPlants(plants []int, capacity int) int {
	// 模拟
	var (
		result int
		remain = capacity
	)
	for i := 0; i < len(plants); i++ {
		if remain >= plants[i] {
			result++
			remain -= plants[i]
		} else {
			remain = capacity - plants[i]
			result += i + i + 1
		}
	}
	return result
}

func wateringPlants2(plants []int, capacity int) int {
	// 模拟
	var (
		result int
		remain = capacity
		i      int
	)

	for i < len(plants) {
		if remain >= plants[i] {
			result++
			remain -= plants[i]
			i++
		}
		remain = capacity
		result += i + i
	}
	return result
}
