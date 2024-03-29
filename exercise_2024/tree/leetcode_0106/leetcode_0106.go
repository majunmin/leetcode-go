package leetcode_0106

import (
	"slices"

	. "github.com/majunmin/leetcode-go/common"
)

//https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/description/

func buildTree(inorder []int, postorder []int) *TreeNode {
	// param check
	if len(inorder) != len(postorder) {
		panic("invalid param")
	}
	if len(inorder) == 0 {
		return nil
	}
	// nodeVal
	nodeVal := postorder[len(postorder)-1]
	node := &TreeNode{Val: nodeVal}
	idxOfInorder := slices.Index(inorder, nodeVal)
	node.Left = buildTree(inorder[:idxOfInorder], postorder[:idxOfInorder])
	node.Right = buildTree(inorder[idxOfInorder+1:], postorder[idxOfInorder:len(postorder)-1])
	return node
}
