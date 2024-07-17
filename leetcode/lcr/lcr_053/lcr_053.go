package lcr_053

import . "github.com/majunmin/leetcode-go/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	//
	stack := make([]*TreeNode, 0)
	var pFlag bool
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			node = node.Left
		}
		// pop stack
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pFlag {
			return node
		}
		if node == p {
			pFlag = true
		}
		// node
		node = node.Right
	}
	return nil
}
