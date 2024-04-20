package leetcode_1026

import (
	. "github.com/majunmin/leetcode-go/common"
)

// https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/?envType=daily-question&envId=2024-04-13
func maxAncestorDiff(root *TreeNode) int {
	// dfs 遍历树的 每个分支，记录最大值和最小值.()
	return dfs(root, root.Val, root.Val)

}

func dfs(node *TreeNode, maxVal int, minVal int) int {
	// terminate
	if node == nil {
		return maxVal - minVal
	}
	var res int
	if node.Left != nil {
		res = max(res, dfs(node.Left, max(maxVal, node.Left.Val), min(minVal, node.Left.Val)))
	}
	if node.Right != nil {
		res = max(res, dfs(node.Right, max(maxVal, node.Right.Val), min(minVal, node.Right.Val)))
	}
	return res
}
