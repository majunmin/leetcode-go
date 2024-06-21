package leetcode_0129

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/sum-root-to-leaf-numbers/?envType=study-plan-v2&envId=top-interview-150
func sumNumbers(root *TreeNode) int {
	// dfs
	var result int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, cur int) {
		// termiante
		if node == nil {
			return
		}
		cur = cur*10 + node.Val
		if node.Left == nil && node.Right == nil {
			result += cur
			return
		}
		dfs(node.Left, cur)
		dfs(node.Right, cur)
	}
	return result
}
