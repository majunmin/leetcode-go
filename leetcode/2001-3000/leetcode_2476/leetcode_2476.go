package leetcode_2476

import . "github.com/majunmin/leetcode-go/common"

func closestNodes(root *TreeNode, queries []int) [][]int {
	// 0. param check
	if root == nil || len(queries) == 0 {
		return nil
	}
	//1. 中序遍历
	var inorder []int
	dfs(root, &inorder)

	//2. findRange
	var result [][]int
	for _, num := range queries {
		left, right := -1, -1
		// 返回左边第一个小于等于 target 的值的索引
		idx := binarySearch(inorder, num)
		if inorder[idx] == num {
			left, right = num, num
		} else {
			if inorder[idx] < num {
				left = inorder[idx]
			}
			if inorder[idx] > num {
				right = idx
			} else if idx < len(inorder)-1 {
				right = inorder[idx+1]
			}
		}
		result = append(result, []int{left, right})
	}
	return result
}

// 返回左边第一个小于等于 target 的值的索引
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)>>1
		if target <= arr[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func dfs(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	dfs(node.Left, result)
	*result = append(*result, node.Val)
	dfs(node.Right, result)
}
