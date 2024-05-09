package leetcode_0105

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildTreeHelper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}

func buildTreeHelper(preorder []int, preStart int, preEnd int, inorder []int, inStart int, inEnd int) *TreeNode {
	if preStart > preEnd || preEnd-preStart != inEnd-inStart {
		return nil
	}
	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}
	rootIdx := findArrayIndex(inorder, rootVal)
	root.Left = buildTreeHelper(preorder, preStart+1, preStart+rootIdx-inStart, inorder, inStart, rootIdx-1)
	root.Right = buildTreeHelper(preorder, preStart+rootIdx-inStart+1, preEnd, inorder, rootIdx+1, inEnd)
	return root
}

func findArrayIndex(nums []int, num int) int {
	for i, n := range nums {
		if n == num {
			return i
		}
	}
	return -1
}

// 采用  截取  array的方式
func recurSolution1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	roorInorderIdx := findArrayIndex(inorder, preorder[0])
	root.Left = buildTree(preorder[1:roorInorderIdx+1], inorder[:roorInorderIdx])
	root.Right = buildTree(preorder[roorInorderIdx+1:], inorder[roorInorderIdx+1:])
	return root
}
