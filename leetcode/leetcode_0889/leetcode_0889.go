package leetcode_0889

import (
	"slices"

	. "github.com/majunmin/leetcode-go/common"
)

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	// param check
	if len(preorder) != len(postorder) || preorder[0] != postorder[len(postorder)-1] {
		panic("invalid param")
	}
	return buildTree(preorder, postorder)
}

func buildTree(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	//
	node := &TreeNode{Val: preorder[0]}
	preorder = preorder[1:]
	postorder = postorder[:len(postorder)-1]
	if len(preorder) > 0 {
		if preorder[0] == postorder[len(postorder)-1] {
			node.Left = buildTree(preorder, postorder)
		} else {
			rootIDx := slices.Index(preorder, postorder[len(postorder)-1])
			node.Left = buildTree(preorder[:rootIDx], postorder[:rootIDx])
			node.Right = buildTree(preorder[rootIDx:], postorder[rootIDx:])
		}
	}
	return node
}
