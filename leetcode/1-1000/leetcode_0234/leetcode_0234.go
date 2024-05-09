package leetcode_0234

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/palindrome-linked-list/?envType=study-plan-v2&envId=top-100-liked
func isPalindrome(head *ListNode) bool {

	return solution3(head)
}

// 1. 快慢指针， 找到链表中间节点
// 2. 链表的还原.  空间复杂度为O(1)
func solution3(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	next := slow.Next
	tail := reverse(next)
	tmp := tail
	slow.Next = nil
	result := true
	for tail != nil && head != nil {
		if tail.Val != head.Val {
			result = false
			break
		}
		head = head.Next
		tail = tail.Next
	}
	slow.Next = reverse(tmp)
	return result
}

func reverse(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}
	tail := reverse(node.Next)
	node.Next.Next = node
	node.Next = nil
	return tail
}

// 1_upbottom 想法比较简单，将链表转化为数组， 利用双指针解法解决

// 递归的方式， 空间复杂度依旧是  O(n)
func solution2(head *ListNode) bool {
	frontPointer := head
	var recursiveCheck func(node *ListNode) bool
	recursiveCheck = func(node *ListNode) bool {
		if node != nil {
			if !recursiveCheck(node.Next) {
				return false
			}
			if frontPointer.Val != node.Val {
				return false
			}
			frontPointer = frontPointer.Next
		}
		return true
	}

	return recursiveCheck(head)
}
