package leetcode_0100

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/same-tree/?envType=study-plan-v2&envId=top-interview-150
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// terminate
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}
