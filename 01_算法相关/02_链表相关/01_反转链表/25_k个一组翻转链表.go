package main

/*
1 -> 2 -> 3 -> 4 -> 5
翻转后:
2 -> 1 -> 4-> 3 -> 5
解题思路:
1. 两两交换链表中的节点: k=2
2. 先计算链表的总长度count
3. 按照k个进行分组;不满足k个就不分组了
*/
// k个一组翻转链表: 统计链表长度 +
func reverseKGroup(head *ListNode, k int) *ListNode {
	// 统计链表的长度
	count := 0
	for cur := head; cur != nil; cur = cur.Next {
		count++
	}

	// k个一组的遍历参数
	dummy := &ListNode{Next: head}
	cur := head
	pre := dummy

	// 链表的反转参数
	var preNode *ListNode
	curNode := head

	// 按k个一组进行划分成多个组;
	for ; count >= k; count -= k {
		// k个节点进行翻转
		for i := 0; i < k; i++ {
			nextNode := curNode.Next
			curNode.Next = preNode
			preNode = curNode
			curNode = nextNode
		}

		// 前面一组的当前节点的next指针指向后面一组的当前节点; 因为后面一组的当前节点已经是反转过来的
		// 1后面挂载是3(curNode) 0后面挂载是2(preNode)
		pre.Next = preNode
		cur.Next = curNode
		pre = cur
		cur = curNode
	}
	return dummy.Next
}
