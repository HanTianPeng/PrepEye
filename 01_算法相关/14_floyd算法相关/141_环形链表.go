package main

/*
解题思路:
1. 龟兔赛跑原理->创建一个快指针、一个慢指针，如果成环，一定会相遇
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	// 创建一个快指针、一个慢指针
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
