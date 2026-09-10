package main

/*
解题思路:
链表 1 -> 2 -> 3 -> 3 -> 3 -> 4 -> 4 -> 5
变成 1 -> 2 -> 5
*/
func deleteDuplicates2(head *ListNode) *ListNode {
	// 创建一个哨兵节点
	dummy := &ListNode{Next: head}
	// 遍历
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		val := cur.Next.Val
		if val == cur.Next.Next.Val {
			for cur.Next != nil && cur.Next.Val == val {
				// 将cur节点的next指针指向下一个不相等的节点
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
