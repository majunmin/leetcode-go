package offer_007

import . "github.com/majunmin/leetcode-go/offer"

// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//
//

// https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/?envType=study-plan&id=lcof
func buildTree(preorder []int, inorder []int) *TreeNode {
	root := buildTreeHelper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
	return root
}

func buildTreeHelper(preorder []int, preLeft int, preRight int, inorder []int, inLeft int, inRight int) *TreeNode {
	// terminate
	if preLeft > preRight {
		return nil
	}

	node := &TreeNode{Val: preorder[preLeft]}
	if preLeft == preRight {
		return node
	}

	idx := findIdx(inorder, preorder[preLeft])
	leftLen := idx - inLeft
	node.Left = buildTreeHelper(preorder, preLeft+1, preLeft+leftLen, inorder, inLeft, idx-1)
	node.Right = buildTreeHelper(preorder, preLeft+leftLen+1, preRight, inorder, idx+1, inRight)
	return node
}

func findIdx(inorder []int, val int) int {
	for i, num := range inorder {
		if num == val {
			return i
		}
	}
	return -1
}
