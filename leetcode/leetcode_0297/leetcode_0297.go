package leetcode_0297

import (
	. "github.com/majunmin/leetcode-go/common"
	"strconv"
	"strings"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (cc *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	var stack []*TreeNode
	stack = append(stack, root)
	var res []string
	for len(stack) > 0 {
		l := len(stack)
		for i := 0; i < l; i++ {
			node := stack[0]
			stack = stack[1:]
			if node == nil {
				res = append(res, "nil")
				continue
			}
			res = append(res, strconv.Itoa(node.Val))
			stack = append(stack, node.Left)
			stack = append(stack, node.Right)
		}
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (cc *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	dat := strings.Split(data, ",")
	var stack []*TreeNode
	if dat[0] == "nil" {
		return nil
	}

	val, _ := strconv.Atoi(dat[0])
	root := &TreeNode{Val: val}
	stack = append(stack, root)

	idx := 1
	for idx < len(dat) {
		node := stack[0]
		stack = stack[1:]
		leftVal, _ := strconv.Atoi(dat[idx])
		rightVal, _ := strconv.Atoi(dat[idx+1])

		if dat[idx] != "nil" {
			leftNode := &TreeNode{Val: leftVal}
			node.Left = leftNode
			stack = append(stack, leftNode)
		}

		if dat[idx+1] != "nil" {
			rightNode := &TreeNode{Val: rightVal}
			node.Right = rightNode
			stack = append(stack, rightNode)
		}

		idx += 2
	}
	return root
}
