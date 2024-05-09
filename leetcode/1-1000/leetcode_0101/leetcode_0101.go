// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-22 21:27
package leetcode_0101

import . "github.com/majunmin/leetcode-go/common"

func isSymmetric(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}
	return isSymmetric(root.Left) && isSymmetric(root.Right) && (root.Left.Val == root.Right.Val)

}
