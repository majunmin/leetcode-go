package leetcode_0590

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/n-ary-tree-postorder-traversal/
func postorder(root *Node) []int {
	//var result []int
	//postorderHelper(root, &result)
	//return result
	return stackSolution(root)
}

func stackSolution(root *Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	var stack []*Node
	var lastNode *Node

	stack = append(stack, root)
	for len(stack) > 0 {
		// peek
		node := stack[len(stack)-1]
		if len(node.Children) == 0 || node.Children[len(node.Children)-1] == lastNode {
			result = append(result, node.Val)
			// 出栈
			stack = stack[:len(stack)-1]
		} else {
			for i := len(node.Children); i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
		}
	}
	return result

}

// 递归解法
func postorderHelper(root *Node, result *[]int) {
	if root == nil {
		return
	}
	// 遍历 所有子节点
	for _, cNode := range root.Children {
		postorderHelper(cNode, result)
	}

	*result = append(*result, root.Val)
}
