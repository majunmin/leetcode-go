package leetcode_0108

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/?envType=study-plan-v2&envId=top-100-liked
func sortedArrayToBST(nums []int) *TreeNode {
	// param check
	if len(nums) == 0 {
		return nil
	}
	return buildTree(nums, 0, len(nums)-1)
}

func buildTree(nums []int, l int, r int) *TreeNode {
	// termiante
	if r < l {
		return nil
	}
	if l == r {
		return &TreeNode{Val: nums[l]}
	}
	// process cur level
	mid := l + (r-l)>>1
	node := &TreeNode{Val: nums[mid]}
	node.Left = buildTree(nums, l, mid-1)
	node.Right = buildTree(nums, mid+1, r)
	return node
}
