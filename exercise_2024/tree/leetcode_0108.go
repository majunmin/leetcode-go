package tree

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/?envType=study-plan-v2&envId=top-100-liked
func sortedArrayToBST(nums []int) *TreeNode {
	// param check
	if len(nums) == 0 {
		return nil
	}
	return buildBST(nums, 0, len(nums)-1)
}

func buildBST(nums []int, l int, r int) *TreeNode {
	// terminate
	if l > r {
		return nil
	}
	if l == r {
		return &TreeNode{Val: nums[l]}
	}
	mid := l + (r-l)>>1
	node := &TreeNode{Val: nums[mid]}
	node.Left = buildBST(nums, l, mid-1)
	node.Right = buildBST(nums, mid+1, r)
	return node
}
