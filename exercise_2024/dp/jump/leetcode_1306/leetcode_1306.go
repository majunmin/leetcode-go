package leetcode_1306

// https://leetcode.cn/problems/jump-game-iii/
func canReach(arr []int, start int) bool {
	//  bfs
	// param check
	n := len(arr)
	if n <= start {
		return false
	}
	//
	queue := make([]int, 0)
	queue = append(queue, start)
	visited := make([]bool, n)
	visited[start] = true
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curIdx := queue[i]
			if arr[curIdx] == 0 {
				return true
			}
			item := arr[curIdx]
			lIdx, rIdx := curIdx+item, curIdx-item
			if lIdx >= 0 && lIdx < len(arr) && !visited[lIdx] {
				queue = append(queue, lIdx)
				visited[lIdx] = true
			}
			if rIdx >= 0 && rIdx < len(arr) && !visited[rIdx] {
				queue = append(queue, rIdx)
				visited[rIdx] = true
			}
		}
		queue = queue[size:]
	}
	return false
}
