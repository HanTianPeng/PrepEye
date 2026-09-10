package main

/*
解题思路:
1. 从题目删除倒数第n个节点就能看到这个题一定是用快慢指针
2. 先让快指针先走n步骤
3. 再快慢指针同时走，直到快指针结束
4. 改变指向;因为需要改变指向,所以需要记录慢指针的pre节点
5. 边界处理,pre节点有可能为nil
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 创建一个哨兵指针
	dummyHead := &ListNode{Next: head}

	// 创建快慢指针
	slow, fast := dummyHead, dummyHead

	// 快指针先走n步骤
	for ; n > 0; n-- {
		fast = fast.Next
	}

	// 快慢指针同时出发,直到快指针到达尾部
	for fast != nil && fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// 改变指针指向
	slow.Next = slow.Next.Next
	return dummyHead.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	// 创建快慢指针
	slow, fast := head, head
	var preNode *ListNode
	// 快指针先走n步骤
	for n-1 > 0 {
		fast = fast.Next
		n--
	}

	// 快慢指针同时走
	for fast.Next != nil {
		preNode = slow
		fast = fast.Next
		slow = slow.Next
	}

	// 链表长度为n，倒数第n个;则相当于剔除头节点
	if preNode == nil {
		return head.Next
	}

	preNode.Next = slow.Next
	return head
}
