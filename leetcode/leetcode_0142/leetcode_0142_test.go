// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-22 22:55
package leetcode_0142

import (
	"fmt"
	"github.com/majunmin/leetcode-go/common"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	node1 := &common.ListNode{
		Val: 1,
	}
	node2 := &common.ListNode{
		Val: 2,
	}
	node3 := &common.ListNode{
		Val: 3,
	}
	node4 := &common.ListNode{
		Val: 4,
	}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	res := detectCycle(node1)
	fmt.Println(res)

}
