package leetcode_0222

import . "github.com/majunmin/leetcode-go/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//https://leetcode.cn/problems/count-complete-tree-nodes/?envType=study-plan-v2&envId=top-interview-150
func countNodes(root *TreeNode) int {
	// terminate
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// 利用题目中的 完全二叉树的 特性
func countNodes2(root *TreeNode) int {
	// terminate
	if root == nil {
		return 0
	}
	lDepth := contLevel(root.Left)
	rDepth := contLevel(root.Right)
	if lDepth == rDepth {
		// 左树是完全的, 右树不完全.
		return countNodes2(root.Right) + 1<<lDepth
	} else {
		return countNodes2(root.Left) + 1<<rDepth
	}
}

func contLevel(node *TreeNode) int {
	var depth int
	for node != nil {
		depth++
		node = node.Left
	}
	return depth
}
