package leetcode_0297

import (
	"fmt"
	"testing"

	"github.com/majunmin/leetcode-go/common"
)

func Test_Tree(t *testing.T) {
	node1 := &common.TreeNode{Val: 1}
	node2 := &common.TreeNode{Val: 2}
	node3 := &common.TreeNode{Val: 3}
	node4 := &common.TreeNode{Val: 4}
	node5 := &common.TreeNode{Val: 5}

	node1.Left = node2
	node1.Right = node3
	node3.Left = node4
	node3.Right = node5

	codec := Constructor()
	fmt.Println(codec.serialize(node1))

}

func Test_TreeDeserilize(t *testing.T) {
	codec := Constructor()
	treeNode := codec.deserialize("1,2,nil,nil,3,4,nil,nil,5,nil,nil")
	fmt.Println(treeNode)
}
