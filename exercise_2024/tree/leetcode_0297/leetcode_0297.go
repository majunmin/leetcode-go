package leetcode_0297

import (
	"strconv"
	"strings"

	. "github.com/majunmin/leetcode-go/common"
)

const (
	NULL = "null"
)

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var result []string
	recursiveSerialize(root, &result)
	return strings.Join(result, ",")
}

func recursiveSerialize(node *TreeNode, result *[]string) {
	if node == nil {
		*result = append(*result, "null")
		return
	}
	val := strconv.Itoa(node.Val)
	*result = append(*result, val)
	recursiveSerialize(node.Left, result)
	recursiveSerialize(node.Right, result)

}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	items := strings.Split(data, ",")
	root := recursiveDserialize(&items)
	return root
}

func recursiveDserialize(items *[]string) *TreeNode {
	if len(*items) == 0 {
		return nil
	}
	curVal := (*items)[0]
	if curVal == "null" {
		return nil
	}
	val, _ := strconv.Atoi(curVal)
	node := &TreeNode{Val: val}
	*items = (*items)[1:]
	node.Left = recursiveDserialize(items)
	*items = (*items)[1:]
	node.Right = recursiveDserialize(items)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
