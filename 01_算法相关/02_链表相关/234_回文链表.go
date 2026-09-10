package main

/*
解题思路:
1 -> 2 -> 2 -> 1 回文链表
1. 先获取中间的节点 -> 快慢双指针
2. 如何判断回文 -> 反转链表
3. 判断回文遍历
*/
func isPalindrome(head *ListNode) bool {
	// 876_链表的中间节点: 快慢指针实现,快指针走2步,慢指针走1步
	mid := middleNode(head)

	// 206_反转链表: 前后指针实现,
	reverse := reverseList(mid)
	// 判断是否为回文链表
	for reverse != nil {
		if reverse.Val != head.Val {
			return false
		}
		reverse = reverse.Next
		head = head.Next
	}
	return true
}
