package leetcode_1103

// https://leetcode.cn/problems/distribute-candies-to-people/
func distributeCandies(candies int, num_people int) []int {
	if num_people <= 0 {
		return nil
	}

	// 模拟
	var (
		result = make([]int, num_people)
		total  int
		curIdx int
	)
	item := 1
	for total < candies {
		if item+total >= candies {
			result[curIdx] += candies - total
			break
		}
		result[curIdx] += item
		total += item
		curIdx = (curIdx + 1) % num_people
		item += 1
	}
	return result
}
