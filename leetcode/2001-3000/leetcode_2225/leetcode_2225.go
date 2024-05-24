package leetcode_2225

// https://leetcode.cn/problems/find-players-with-zero-or-one-losses/
func findWinners(matches [][]int) [][]int {
	var (
		lossCnt = make(map[int]int)
		result  = make([][]int, 2)
	)
	for _, item := range matches {
		if _, exist := lossCnt[item[0]]; !exist {
			lossCnt[item[0]] = 0
		}
		lossCnt[item[1]]++
	}
	for id, cnt := range lossCnt {
		if cnt < 2 {
			result[cnt] = append(result[cnt], id)
		}
	}
	return result
}
