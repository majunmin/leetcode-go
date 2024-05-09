package leetcode_2385

import . "github.com/majunmin/leetcode-go/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
	return solution4(root, start)
}

func solution4(node *TreeNode, start int) int {

	var (
		ans int
		dfs func(*TreeNode) int
	)
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lDepth, rDepth := dfs(node.Left), dfs(node.Right)
		// find startNode
		if node.Val == start {
			ans = 1 - min(lDepth, rDepth)
			return 1
		}
		if lDepth > 0 || rDepth > 0 {
			// 更新 ans
			ans = max(ans, abs(lDepth)+abs(rDepth)+1)
			return max(lDepth, rDepth) + 1 // 自动取正数
		}

		return min(lDepth, rDepth) - 1
	}
	dfs(node)
	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// 3.1 一次遍历的解法,类似于  计算树的直径
func solution3(root *TreeNode, start int) int {
	var maxDepth func(*TreeNode, int) (int, bool)
	var ans int
	maxDepth = func(node *TreeNode, start int) (int, bool) {
		if node == nil {
			return 0, false
		}
		lDepth, lFound := maxDepth(node.Left, start)
		rDepth, rFound := maxDepth(node.Right, start)
		if node.Val == start {
			// 如果当前节点是  start 节点
			// 更新答案
			ans = max(lDepth, rDepth)
			// todo
			return 1, true
		}
		if lFound || rFound {
			// 更新ans
			ans = max(ans, lDepth+rDepth)
			if lFound {
				return lDepth + 1, true
			}
			return rDepth + 1, true
		}
		return max(lDepth, rDepth) + 1, false

	}
	maxDepth(root, start)
	return ans
}

// 2. 有父节点指向子节点的指针已经存在， 创建 子节点指向父节点的指针(map)
func solution2(root *TreeNode, start int) int {
	var (
		from      = make(map[*TreeNode]*TreeNode)
		startNode *TreeNode
	)

	var buildParent func(node, parent *TreeNode, parents map[*TreeNode]*TreeNode)
	buildParent = func(node, parent *TreeNode, parents map[*TreeNode]*TreeNode) {
		if node == nil {
			return
		}
		if node.Val == start {
			startNode = node
		}
		if parent != nil {
			parents[node] = parent
		}
		buildParent(node.Left, node, parents)
		buildParent(node.Right, node, parents)
	}
	buildParent(root, nil, from)

	// find 最大直径
	return calcDiameter(startNode, nil, from)
}

func calcDiameter(node *TreeNode, parent *TreeNode, from map[*TreeNode]*TreeNode) int {
	if node == nil {
		return -1
	}
	var res int
	if node.Left != parent {
		res = max(res, calcDiameter(node.Left, node, from))
	}
	if node.Right != parent {
		res = max(res, calcDiameter(node.Right, node, from))
	}
	if from[node] != parent {
		res = max(res, calcDiameter(from[node], node, from))
	}
	return res + 1
}

// 1. 先构建图, 之后在dfs 遍历
func solution1(root *TreeNode, start int) int {
	// 将二叉树 转化为图, 之后进行 bfs 传染.
	var (
		adj     = make(map[int][]int, 0)
		queue   = make([]int, 0)
		visited = make(map[int]bool)
		result  = -1
	)
	dfs(root, adj)
	// bfs
	// 避免重复遍历
	queue = append(queue, start)
	visited[start] = true
	for len(queue) > 0 {
		result++
		size := len(queue)
		for i := 0; i < size; i++ {
			if visited[queue[i]] {
				continue
			}
			visited[queue[i]] = true
			for _, next := range adj[queue[i]] {
				queue = append(queue, next)
			}
		}
		queue = queue[size:]
	}
	return -1
}

func dfs(root *TreeNode, adj map[int][]int) {
	if root.Left != nil {
		adj[root.Val] = append(adj[root.Val], root.Left.Val)
		adj[root.Left.Val] = append(adj[root.Left.Val], root.Val)
		dfs(root.Left, adj)
	}
	if root.Right != nil {
		adj[root.Val] = append(adj[root.Val], root.Right.Val)
		adj[root.Right.Val] = append(adj[root.Right.Val], root.Val)
		dfs(root.Right, adj)
	}
}
