package leetcode_0210

// https://leetcode.cn/problems/course-schedule-ii/
func findOrder(numCourses int, prerequisites [][]int) []int {
	//拓扑排序 类似 与 题目(207)
	var (
		adj      = make([][]int, numCourses)
		inDegree = make([]int, numCourses)
		learnCnt int
	)

	for _, dep := range prerequisites {
		adj[dep[1]] = append(adj[dep[1]], dep[0])
		inDegree[dep[0]]++
	}

	queue := make([]int, 0)
	for i, d := range inDegree {
		if d == 0 {
			queue = append(queue, i)
			learnCnt++
		}
	}
	for i := 0; i < len(queue); i++ {
		item := queue[i]
		for _, next := range adj[item] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
				learnCnt++
			}
		}
	}
	if learnCnt < numCourses {
		return nil
	}
	return queue
}
