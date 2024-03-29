package leetcode_0105

import . "github.com/majunmin/leetcode-go/common"

func buildTree(preorder []int, inorder []int) *TreeNode {
	// param check
	if len(preorder) == 0 {
		return nil
	}
	//
	curNodeVal := preorder[0]
	inorderIdx := findIdx(inorder, curNodeVal)
	node := &TreeNode{Val: curNodeVal}
	node.Left = buildTree(preorder[1:inorderIdx+1], inorder[:inorderIdx])
	node.Right = buildTree(preorder[inorderIdx+1:], inorder[inorderIdx+1:])
	return node
}

func findIdx(arr []int, target int) int {
	for i := range arr {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
