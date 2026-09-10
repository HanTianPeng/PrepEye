package main

func sortList(head *ListNode) *ListNode {
	// 边界判断: 如果链表为null或则只有一个节点则无需排序
	if head == nil || head.Next == nil {
		return head
	}
	// 二分法;获取中间节点;并将链表一分为2(head->mid) (mid->tail)
	head2 := middleNode2(head)
	// 分而治之
	head = sortList(head)
	head2 = sortList(head2)
	// 合并
	return mergeTwoSortedLists(head, head2)
}

// 求出链表的中间节点
func middleNode2(head *ListNode) *ListNode {
	// 通过快慢指针;快指针走两步;慢指针走一步
	preNode, fast, slow := head, head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		preNode = slow
		slow = slow.Next
	}
	preNode.Next = nil
	return slow
}

func mergeTwoSortedLists(list1, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		// node1小于node2
		list1.Next = mergeTwoSortedLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoSortedLists(list1, list2.Next)
	return list2
}
