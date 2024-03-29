package leetcode_0124

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/binary-tree-maximum-path-sum/?envType=study-plan-v2&envId=top-100-liked
func maxPathSum(root *TreeNode) int {
	res, _ := maxPathSumHelper(root)
	return res
}

func maxPathSumHelper(node *TreeNode) (int, int) {
	if node == nil {
		return 0, 0
	}
	//
	lres, lr := maxPathSumHelper(node.Left)
	rres, rr := maxPathSumHelper(node.Right)
	return max(lres, rres, lr+rr+node.Val), max(lr, rr) + node.Val
}
