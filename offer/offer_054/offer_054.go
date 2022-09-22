package offer_054

import . "github.com/majunmin/leetcode-go/offer"

//https://leetcode.cn/problems/er-cha-sou-suo-shu-de-di-kda-jie-dian-lcof/?envType=study-plan&id=lcof
func kthLargest(root *TreeNode, k int) int {

	// param check
	if root == nil {
		return -1
	}

	item := root
	stack := make([]*TreeNode, 0)
	rank := 0
	for len(stack) > 0 || item != nil {
		for item != nil {
			stack = append(stack, item)
			item = item.Right
		}
		// pop item
		item = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		rank++
		if rank == k {
			return item.Val
		}
		if item.Left != nil {
			item = item.Left
			continue
		}
		item = nil
	}

	return -1
}
