package leetcode_0106

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/
func buildTree(inorder []int, postorder []int) *TreeNode {
	// param check
	if len(inorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}

	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}
	// recursive
	rIdx := findIndexInArray(inorder, rootVal)
	root.Left = buildTree(inorder[0:rIdx], postorder[:rIdx])
	root.Right = buildTree(inorder[rIdx+1:], postorder[rIdx:len(postorder)-1])
	return root
}

func findIndexInArray(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}
