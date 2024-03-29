package leetcode_0114

import . "github.com/majunmin/leetcode-go/common"

func flatten(root *TreeNode) {
	// param check
	flattenHelp(root)

}

func flattenHelp(node *TreeNode) *TreeNode {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return node
	}

	var leftEnd, rightEnd *TreeNode
	rightEnd = flattenHelp(node.Right)
	if node.Left != nil {
		leftEnd = flattenHelp(node.Left)
		leftEnd.Right = node.Right
		node.Right = node.Left
	}
	if rightEnd != nil {
		return rightEnd
	}
	return leftEnd
}
