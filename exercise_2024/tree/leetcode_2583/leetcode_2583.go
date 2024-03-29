package leetcode_2583

import (
	"slices"

	. "github.com/majunmin/leetcode-go/common"
)

// https://leetcode.cn/problems/kth-largest-sum-in-a-binary-tree/
func kthLargestLevelSum(root *TreeNode, k int) int64 {
	if root == nil {
		return -1
	}
	var levelSum []int64
	// bfs
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		var curSum int64
		for i := 0; i < size; i++ {
			item := queue[i]
			curSum += int64(item.Val)
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
		}
		queue = queue[size:]
		levelSum = append(levelSum, curSum)
	}
	if len(levelSum) < k {
		return -1
	}
	slices.Sort(levelSum)
	return levelSum[k-1]
}
