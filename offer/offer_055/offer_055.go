package offer_055

import . "github.com/majunmin/leetcode-go/offer"

// https://leetcode.cn/problems/er-cha-shu-de-shen-du-lcof/?envType=study-plan&id=lcof
func maxDepth(root *TreeNode) int {
	// terminate
	if root == nil {
		return 0
	}
	//
	return 1 + maxInt(maxDepth(root.Left), maxDepth(root.Right))
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}
