package leetcode_0236

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return solution2(root, p, q)
}

// 方法二
// 1. 通过一遍dfs遍历, 用一个 hash表 存储所有节点的父节点.
// 2. 建立两条路径 p -> root 和 q -> root
// 3. 两条路径的最早的交点就是  最近公共祖先节点
func solution2(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	var (
		parent  = make(map[*TreeNode]*TreeNode)
		dfs     func(*TreeNode, *TreeNode)
		visited = make(map[*TreeNode]bool)
	)

	dfs = func(node, pre *TreeNode) {
		if node == nil {
			return
		}
		parent[node] = pre
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)
	// 2. 构建路径
	for p != nil {
		visited[p] = true
		p = parent[p]
	}
	for q != nil {
		if visited[q] {
			return q
		}
		q = parent[q]
	}
	return nil
}

// 1. 递归解法
func solution1(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	// terminate
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// process
	lson := lowestCommonAncestor(root.Left, p, q)
	rson := lowestCommonAncestor(root.Right, p, q)
	if lson != nil && rson != nil {
		return root
	}
	if lson == nil {
		return rson
	}
	return lson
}
