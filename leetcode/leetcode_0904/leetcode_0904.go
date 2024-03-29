package leetcode_0904

// https://leetcode.cn/problems/fruit-into-baskets/
func totalFruit(fruits []int) int {
	// param check
	if len(fruits) < 2 {
		return len(fruits)
	}
	var (
		needs     int // 窗口内水果种类数
		result    int // 篮子中的最大水果数
		cnt       int // 当前篮子中水果数
		fruitCnts = make(map[int]int)
		l, r      int
	)
	for ; r < len(fruits); r++ {
		fruitCnts[fruits[r]]++
		cnt++
		if fruitCnts[fruits[r]] == 1 {
			needs++
		}
		if needs <= 2 {
			result = max(result, cnt)
			continue
		}
		// 窗口内 水果类型 > 2
		// shrink window
		for l < r {
			fruitCnts[fruits[l]]--
			cnt--
			if fruitCnts[fruits[l]] == 0 {
				l++
				needs--
				break
			}
			l++
		}
	}
	return result
}
