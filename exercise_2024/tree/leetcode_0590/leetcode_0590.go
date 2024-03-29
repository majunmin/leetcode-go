package leetcode_0590

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	return recursiveSolution(root)
}

func recursiveSolution(root *Node) []int {
	if root == nil {
		return nil
	}
	var result []int
	recursiveSolutionInner(root, &result)
	return result
}

func recursiveSolutionInner(root *Node, result *[]int) {
	// terminate
	if root == nil {
		return
	}
	for i := len(root.Children) - 1; i >= 0; i-- {
		recursiveSolutionInner(root.Children[i], result)
	}
	*result = append(*result, root.Val)
}

// 利用栈的迭代解法
func iterSolution(root *Node) []int {
	// param check
	if root == nil {
		return nil
	}
	stack := make([]*Node, 0)
	stack = append(stack, root)
	var result []int
	var prev *Node
	for len(stack) > 0 {
		item := stack[len(stack)-1]
		if len(item.Children) == 0 || prev == item.Children[len(item.Children)-1] {
			stack = stack[:len(stack)-1]
			result = append(result, item.Val)
			prev = item
		} else {
			for i := len(item.Children) - 1; i >= 0; i-- {
				stack = append(stack, item.Children[i])
			}
		}
	}
	return result
}
