package exercise_2024

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/?envType=study-plan-v2&envId=top-100-liked
func buildTree(preorder []int, inorder []int) *TreeNode {
	// param check
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	node := &TreeNode{Val: rootVal}
	idx := findIdx(inorder, rootVal)
	node.Left = buildTree(preorder[1:idx+1], inorder[0:idx])
	node.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return node
}

// 一定能找到
func findIdx(inorder []int, val int) int {
	for i := range inorder {
		if inorder[i] == val {
			return i
		}
	}
	return -1
}
