package lcr_050

import (
	. "github.com/majunmin/leetcode-go/common"
)

//https://leetcode.cn/problems/6eUYwP/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) int {
	// terminate
	if root == nil {
		return 0
	}
	cnt := countPaths(root, targetSum)
	cnt += pathSum(root.Left, targetSum-root.Val)
	cnt += pathSum(root.Right, targetSum-root.Val)
	return cnt
}

func countPaths(node *TreeNode, target int) int {
	if node == nil {
		return 0
	}
	var cnt int
	nextVal := target - node.Val
	if nextVal == 0 {
		cnt = 1
	}
	cnt += countPaths(node.Left, nextVal)
	cnt += countPaths(node.Right, nextVal)
	return cnt
}
