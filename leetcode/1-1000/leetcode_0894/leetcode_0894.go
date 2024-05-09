package leetcode_0894

import . "github.com/majunmin/leetcode-go/common"

func node() *TreeNode {
	return &TreeNode{}
}

// https://leetcode.cn/problems/all-possible-full-binary-trees/
func allPossibleFBT(n int) []*TreeNode {
	var result []*TreeNode
	if n&1 == 0 {
		return result
	}
	if n == 1 {
		return []*TreeNode{node()}
	}
	for i := 1; i < n; i += 2 {
		leftSubTrees := allPossibleFBT(i)
		rightSubTrees := allPossibleFBT(n - i - 1)
		//构建 tree
		for _, lTree := range leftSubTrees {
			for _, rTree := range rightSubTrees {
				root := node()
				root.Left, root.Right = lTree, rTree
				result = append(result, root)
			}
		}
	}
	return result
}

func dpSolution(n int) []*TreeNode {
	// 偶数个节点 不可能构成 FBT
	if n&1 == 0 {
		return nil
	}
	dp := make([][]*TreeNode, n+1)
	dp[1] = []*TreeNode{node()}
	for i := 3; i <= n; i += 2 {
		for j := 1; j < i; j += 2 {
			for _, lTree := range dp[j] {
				for _, rTree := range dp[i-j-1] {
					root := node()
					root.Left, root.Right = lTree, rTree
					dp[i] = append(dp[i])
				}
			}
		}
	}
	return dp[n]
}
