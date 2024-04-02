package leetcode_0331

import "strings"

// https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/?envType=daily-question&envId=2024-03-31
func isValidSerialization(preorder string) bool {
	return inoutDegree(preorder)
}

// https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/?envType=daily-question&envId=2024-03-31
// 利用节点的出度数 == 入度数
// '#'     节点 in = 1, out = 0
// 'digit' 节点 in = 1, out = 2
// root 节点 in = 0 , out = 2
func inoutDegree(preorder string) bool {
	diff := 1
	items := strings.Split(preorder, ",")
	for _, item := range items {
		diff -= 1
		if diff < 0 {
			return false
		}
		if item != "#" {
			diff += 2
		}
	}
	return diff == 0
}

func stackSolution(preorder string) bool {
	// param check
	if len(preorder) < 3 {
		return false
	}
	//
	stack := make([]string, 0)
	items := strings.Split(preorder, ",")
	for i := 0; i < len(items); i++ {
		item := strings.TrimSpace(items[i])
		stack = append(stack, item)
		for len(stack) >= 3 && stack[len(stack)-3] != "#" && stack[len(stack)-1] == "#" && stack[len(stack)-2] == "#" {
			// pop 3次
			stack = stack[:len(stack)-3]
			stack = append(stack, "#")
		}
	}
	return len(stack) == 1 && stack[0] == "#"
}
