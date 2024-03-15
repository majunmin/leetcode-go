package leetcode_2581

// https://leetcode.cn/problems/count-number-of-possible-root-nodes/?envType=daily-question&envId=2024-03-01
func rootCount(edges [][]int, guesses [][]int, k int) int {
	// param check
	if len(guesses) < k {
		return 0
	}
	var (
		adj = make([][]int, len(edges)+1)
		set = make(map[int]bool)
	)
	//1. build graph & set(guess 父子关系)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	for _, e := range guesses {
		set[buildIndex(e[0], e[1])] = true
	}
	// 2.  dfs
	cnt := dfs(adj, set, 0, -1)
	// 3. rootChangeDfs
	res := rootChangeDfs(adj, set, cnt, k, 0, -1)
	// return
	return res
}

func rootChangeDfs(adj [][]int, set map[int]bool, cnt int, k int, cur int, parent int) int {
	var res int
	if cnt >= k {
		res++
	}
	for _, c := range adj[cur] {
		if c == parent {
			continue
		}
		count := cnt
		// 反转当前的 父子关系
		if set[buildIndex(cur, c)] { // 原来是对的, 现在错了
			count--
		}
		if set[buildIndex(c, cur)] { // 原来是错的, 现在对了
			count++
		}
		res += rootChangeDfs(adj, set, count, k, c, cur)
	}
	return res
}

func dfs(adj [][]int, set map[int]bool, cur int, parent int) int {
	// terminate
	var cnt int
	for _, c := range adj[cur] {
		// 无向图不能重复
		if c == parent {
			continue
		}
		if set[buildIndex(cur, c)] {
			cnt++
		}
		cnt += dfs(adj, set, c, cur)
	}
	return cnt
}

func buildIndex(x, y int) int {
	return x<<32 + y
}
