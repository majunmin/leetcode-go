package offer_037

import (
	. "github.com/majunmin/leetcode-go/offer"
	"strconv"
	"strings"
)

var (
	Prefix = "["
	Suffix = "]"
)

// 二叉树的层次遍历
type Codec struct {
}

func (cc *Codec) Serialize(root *TreeNode) string {
	if root == nil {
		return Prefix + "" + Suffix
	}

	dataArr := make([]string, 0)

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			item := queue[i]
			if item == nil {
				dataArr = append(dataArr, "nil")
				continue
			}
			dataArr = append(dataArr, strconv.Itoa(item.Val))
			queue = append(queue, item.Left)
			queue = append(queue, item.Right)
		}
		queue = queue[l:]
	}
	return Prefix + strings.Join(dataArr, ",") + Suffix
}

func (cc *Codec) Deserialize(root string) *TreeNode {
	dataArr, _ := extraData(root)
	if len(dataArr) == 0 {
		return nil
	}

	rootNode := &TreeNode{Val: dataArr[0]}
	queue := make([]*TreeNode, 0)
	queue = append(queue, rootNode)
	idx := 1
	for len(queue) > 0 && idx < len(dataArr) {
		l := len(queue)

		for i := 0; i < l; i++ {
			item := queue[i]
			if dataArr[idx] != -1 {
				item.Left = &TreeNode{Val: dataArr[idx]}
				queue = append(queue, item.Left)
			}
			idx++
			if dataArr[idx] != -1 {
				item.Right = &TreeNode{Val: dataArr[idx]}
				queue = append(queue, item.Right)
			}
			idx++
		}

		queue = queue[l:]
	}
	return rootNode
}

func extraData(root string) ([]int, error) {
	root = strings.TrimSpace(root)
	root = strings.TrimPrefix(root, Prefix)
	root = strings.TrimSuffix(root, Suffix)

	splitStrs := strings.Split(root, ",")
	result := make([]int, len(splitStrs))
	for i := 0; i < len(splitStrs); i++ {
		if strings.EqualFold("nil", splitStrs[i]) {
			result[i] = -1
			continue
		}
		val, err := strconv.ParseInt(splitStrs[i], 10, 64)
		if err != nil {
			return nil, err
		}
		result[i] = int(val)
	}
	return result, nil
}
