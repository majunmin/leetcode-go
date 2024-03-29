package exercise_2024

func canFinish(numCourses int, prerequisites [][]int) bool {
	// param check
	if len(prerequisites) == 0 {
		return true
	}

	var (
		adj      = make([][]int, numCourses)
		inDegree = make([]int, numCourses)
		learnCnt int
	)

	for _, dep := range prerequisites {
		adj[dep[1]] = append(adj[dep[1]], dep[0])
		inDegree[dep[0]]++
	}
	// 遍历, 将 degree == 0 的值加入 queue
	queue := make([]int, 0, numCourses)
	for i, item := range inDegree {
		if item == 0 {
			learnCnt++
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			for _, next := range adj[item] {
				inDegree[next]--
				if inDegree[next] == 0 {
					learnCnt++
					queue = append(queue, next)
				}
			}
		}
		queue = queue[size:]
	}
	return learnCnt == numCourses
}
