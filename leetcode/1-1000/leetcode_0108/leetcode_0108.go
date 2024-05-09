package leetcode_0108

import . "github.com/majunmin/leetcode-go/common"

type pair struct {
	l, r int
}

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
func sortedArrayToBST(nums []int) *TreeNode {
	// param check
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{}
	nodeQueue := make([]*TreeNode, 0)
	rangeQueue := make([]pair, 0)
	nodeQueue = append(nodeQueue, root)
	rangeQueue = append(rangeQueue, pair{0, len(nums) - 1})
	for len(nodeQueue) > 0 {
		cur := nodeQueue[0]
		nodeQueue = nodeQueue[1:]
		p := rangeQueue[0]
		rangeQueue = rangeQueue[1:]
		l, r := p.l, p.r
		mid := l + (r-l)>>1
		cur.Val = nums[mid]
		if l <= mid-1 {
			cur.Left = &TreeNode{}
			nodeQueue = append(nodeQueue, cur.Left)
			rangeQueue = append(rangeQueue, pair{l, mid - 1})
		}
		if mid+1 <= r {
			cur.Right = &TreeNode{}
			nodeQueue = append(nodeQueue, cur.Right)
			rangeQueue = append(rangeQueue, pair{mid + 1, r})
		}
	}
	return root
}

// 解法一: 递归解法
// 1. 使用 左闭右闭区间.
func solution1(nums []int) *TreeNode {
	// param check
	if len(nums) == 0 {
		return nil
	}
	l, r := 0, len(nums)-1
	return genBST(nums, l, r)
}

func genBST(nums []int, l int, r int) *TreeNode {
	if l > r {
		return nil
	}
	if l == r {
		return &TreeNode{Val: nums[l]}
	}
	mid := l + (r-l)>>1
	node := &TreeNode{
		Val: nums[mid],
	}
	node.Left = genBST(nums, l, mid-1)
	node.Right = genBST(nums, mid+1, r)
	return node
}
