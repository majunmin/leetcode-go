package leetcode_1766

//预处理

var (
	mx = 51
	// 互质数字典
	coprimes = make([][]int, mx)
)

// 1. 预计算
func init() {
	for i := 0; i < mx; i++ {
		for j := 0; j < mx; j++ {
			if gcd(i, j) == 1 {
				coprimes[i] = append(coprimes[i], j)
			}
		}
	}
}

// 这个题目需要一些审题技巧.
//
//	1 <= nums[i] <= 50
//	1 <= n <= 105
//
// -  如果遍历每个节点, 在遍历其父节点, 找到最短路径的互质数, 那么时间复杂度是 O(n^2) 的, 一定会超时的.
// 换个思路的话,遍历数字, 记录每个数字出现的depth,
// https://leetcode.cn/problems/tree-of-coprimes/?envType=daily-question&envId=2024-04-11
func getCoprimes(nums []int, edges [][]int) []int {
	//2. 构建无向图
	adj := make(map[int][]int, 0)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	// 构建parent map
	var (
		n     = len(nums)
		depth = make([]int, n)
		pos   = make([]int, n) //存储 num 到 index 的映射关系.
		ans   = make([]int, n)
	)
	// fill default
	for i := 0; i < n; i++ {
		ans[i] = -1
	}
	for i := 0; i < mx; i++ {
		pos[i] = -1
	}
	var dfs func(nums []int, curIdx int, parent int)
	dfs = func(nums []int, curIdx int, parent int) {
		cur := nums[curIdx]
		// 3 . 遍历 nums[curIdx] 的 所有互质数 ( < 50)
		for _, v := range coprimes[cur] {
			// 还没有被访问过
			if pos[v] == -1 {
				continue
			}
			// 选取深度最深的 && 已经访问过的节点就是里当前节点最近的父节点
			if ans[curIdx] == -1 || depth[ans[curIdx]] < depth[pos[v]] {
				ans[curIdx] = pos[v]
			}
		}
		//
		temp := pos[cur]
		// cur在路径上最后一次出现的位置
		pos[cur] = curIdx
		// 4. 遍历子节点
		for _, c := range adj[curIdx] {
			// 避免重复遍历(因为是无向图)
			if c == parent {
				continue
			}
			depth[c] = depth[curIdx] + 1
			dfs(nums, c, curIdx)
		}
		// 回溯
		pos[cur] = temp
	}
	dfs(nums, 0, -1)
	return ans
}

func gcd(a, b int) int {
	// param check
	if b == 0 {
		return a
	}
	// 循环取余
	for a%b != 0 {
		a = a % b
		a, b = b, a
	}
	return b
}
