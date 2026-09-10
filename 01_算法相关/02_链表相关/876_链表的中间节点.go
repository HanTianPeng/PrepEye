package main

/*
解题思路:
链表的中间节点，标准思路: 快慢指针,快指针走2步，慢指针走1步；直到快指针走完结束
*/
func middleNode(head *ListNode) *ListNode {
	// 创建两个指针,快指针,慢指针;快指针走两步,慢指针走一步
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
