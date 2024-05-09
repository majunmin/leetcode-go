package leetcode_0207

// https://leetcode.cn/problems/course-schedule/?envType=study-plan-v2&envId=top-100-liked
func canFinish(numCourses int, prerequisites [][]int) bool {
	return dfsSolution(numCourses, prerequisites)
}

// 拓扑排序实现算法 - DFS遍历实现
func dfsSolution(courses int, prerequisites [][]int) bool {
	// 构建有向图
	// s -> t, 表示 t的执行依赖于 s,(s 先于 t 执行)
	adj := make([][]int, courses)
	// 初始化有向图
	for _, item := range prerequisites {
		adj[item[1]] = append(adj[item[1]], item[0])
	}
	// 标记数组
	// 1： 表示正在遍历过程中, 再次遍历到表示存在环 2：表示正在遍历过程中
	mark := make([]int, courses)
	for i := 0; i < courses; i++ {
		if dfs(i, adj, mark) {
			return false
		}
	}
	return true
}

// @param mark: 标记数组,表示节点在遍历过程中是否访问过， 1: 表示访问过, 2
// @return bool true:图中存在环 false: 不存在环
func dfs(i int, adj [][]int, mark []int) bool {
	// 如果访问过了，就不用再访问了
	if mark[i] == 1 {
		// 从正在访问中，到正在访问中，表示遇到了环
		return true
	}
	if mark[i] == 2 {
		// 表示在访问的过程中没有遇到环，这个节点访问过了
		return false
	}
	// 此时 mark[i] == 0
	mark[i] = 1
	for j := range adj[i] {
		if dfs(adj[i][j], adj, mark) {
			return true
		}
	}
	// i 的所有后继结点都访问完了，都没有存在环，则这个结点就可以被标记为已经访问结束
	// 状态设置为 2
	mark[i] = 2
	// false 表示图中是否存在环
	return false
}

// 拓扑排序的实现算法 - khan
func khanSolution(courses int, prerequisites [][]int) bool {
	// if 课程A 先于 课程B执行(B 依赖 A), 那么  A -> B.
	// 1. 构建邻接表
	adj := make([][]int, courses)
	for i := 0; i < courses; i++ {
		adj[i] = make([]int, 0, courses)
	}
	inDegree := make([]int, courses) // 统计每个节点的 入度(统计每个节点依赖的课程数)
	for _, item := range prerequisites {
		adj[item[1]] = append(adj[item[1]], item[0])
		inDegree[item[0]]++
	}
	queue := make([]int, 0)
	for i := 0; i < courses; i++ {
		// 课程 i 不依赖任何课程, 可以直接学习
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 学习的课程计数
	var cnt int
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		cnt++
		for i := 0; i < len(adj[item]); i++ {
			inDegree[adj[item][i]]--
			if inDegree[adj[item][i]] == 0 {
				queue = append(queue, adj[item][i])
			}
		}
	}
	return cnt == courses
}
