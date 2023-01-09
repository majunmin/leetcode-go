package offer_034

import . "github.com/majunmin/leetcode-go/offer"

// https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/?envType=study-plan&id=lcof
func pathSum(root *TreeNode, target int) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	dfs(&result, root, []int{root.Val}, root.Val, target)
	return result
}

func dfs(result *[][]int, node *TreeNode, path []int, sum, target int) {
	// 剪枝
	if sum > target {
		return
	}

	// terminate
	if node.Left == nil && node.Right == nil && sum == target {
		dst := make([]int, len(path))
		copy(dst, path)
		*result = append(*result, dst)
		return
	}

	//
	if node.Left != nil {
		path = append(path, node.Left.Val)
		dfs(result, node.Left, path, sum+node.Left.Val, target)
		path = path[:len(path)-1]
	}
	//
	if node.Right != nil {
		path = append(path, node.Right.Val)
		dfs(result, node.Right, path, sum+node.Right.Val, target)
		path = path[:len(path)-1]
	}
}
