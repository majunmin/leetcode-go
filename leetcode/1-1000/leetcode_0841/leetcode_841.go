package leetcode_0841

// https://leetcode.cn/problems/keys-and-rooms/
func canVisitAllRooms(rooms [][]int) bool {
	var (
		n       = len(rooms)
		roomNum = n
		visited = make([]bool, n)
		dfs     func(node int)
	)
	dfs = func(node int) {
		// terminate
		if visited[node] {
			return
		}
		visited[node] = true
		roomNum--
		// dfs
		for _, next := range rooms[node] {
			dfs(next)
		}
	}
	dfs(0)
	return roomNum == 0
}

// bfs solution
func canVisitAllRooms2(rooms [][]int) bool {
	var (
		n       = len(rooms)
		roomNum = n
		visited = make([]bool, n)
		queue   = make([]int, 0)
	)
	queue = append(queue, 0)
	visited[0] = true
	roomNum--
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			for _, next := range rooms[item] {
				if visited[next] {
					continue
				}
				queue = append(queue, next)
				visited[next] = true
				roomNum--
			}
		}
		queue = queue[size:]
	}
	return roomNum == 0
}
