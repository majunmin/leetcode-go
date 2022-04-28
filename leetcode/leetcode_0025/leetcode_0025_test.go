// Created By: junmin.ma
// Description: <description>
// Date: 2022-04-26 08:34
package leetcode_0025

import (
	"fmt"
	. "github.com/majunmin/leetcode-go/common"
	"testing"
)

func Test_ReverseKGroupList(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 3}
	node4 := &ListNode{Val: 4}
	node5 := &ListNode{Val: 5}
	node6 := &ListNode{Val: 6}
	node7 := &ListNode{Val: 7}
	node8 := &ListNode{Val: 8}
	node9 := &ListNode{Val: 9}
	node10 := &ListNode{Val: 10}
	node11 := &ListNode{Val: 11}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	node7.Next = node8
	node8.Next = node9
	node9.Next = node10
	node10.Next = node11

	node := reverseKGroup(node1, 11)

	printList(node)
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
}
