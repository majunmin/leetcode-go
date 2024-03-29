package exercise_2024

import . "github.com/majunmin/leetcode-go/common"

// https://leetcode.cn/problems/palindrome-linked-list/?envType=study-plan-v2&envId=top-100-liked
func isPalindrome(head *ListNode) bool {
	frontNode := head
	var dfs func(*ListNode) bool
	dfs = func(node *ListNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Next) {
			return false
		}
		if frontNode.Val != node.Val {
			return false
		}
		frontNode = frontNode.Next
		return true
	}
	return dfs(head)

}
