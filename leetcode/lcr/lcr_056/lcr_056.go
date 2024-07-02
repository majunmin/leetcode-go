package lcr_056

import (
	. "github.com/majunmin/leetcode-go/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//https://leetcode.cn/problems/opLdQZ/
func findTarget(root *TreeNode, k int) bool {
	// hash表 + dfs
	var (
		cache = make(map[int]bool)
	)
	return dfs(root, k, cache)

}

func dfs(node *TreeNode, k int, cache map[int]bool) bool {
	if node == nil {
		return false
	}
	if cache[k-node.Val] {
		return true
	}
	cache[node.Val] = true
	return dfs(node.Left, k, cache) || dfs(node.Right, k, cache)
}
