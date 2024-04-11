package leetcode_2857

// https://leetcode.cn/problems/count-pairs-of-points-with-distance-k/
func countPairs(coordinates [][]int, k int) int {

	var (
		cnts   = make(map[int]int)
		result int
	)

	for _, p := range coordinates {
		x, y := p[0], p[1]
		for i := 0; i < k; i++ {
			result += cnts[cnts[x^i]*2000000+y^(k-i)]
		}
		cnts[cnts[x]*2000000+y] += 1
	}
	return result
}
