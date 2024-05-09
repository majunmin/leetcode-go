// Created By: junmin.ma
// Description: <description>
// Date: 2022-02-21 20:52
package leetcode_0024

import (
	"fmt"
	"testing"

	"github.com/majunmin/leetcode-go/common"
)

func TestSwapPairs(t *testing.T) {
	listNode := &common.ListNode{
		Val: 1,
		Next: &common.ListNode{
			Val: 2,
			Next: &common.ListNode{
				Val: 3,
				Next: &common.ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}

	listNode = swapPairs(listNode)

	fmt.Println(listNode)
	fmt.Println(listNode)
}
