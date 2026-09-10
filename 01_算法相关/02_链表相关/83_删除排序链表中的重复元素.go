package main

/*
解题思路：
1 -> 1 -> 2
变成 1 -> 2
所以第一个节点永远不会被删除,从第二个节点开始才可能被删除;
因此不需要创建哨兵节点,并且遍历从第二个节点开始遍历
*/
func deleteDuplicates(head *ListNode) *ListNode {
	// 由于需要从第二个节点开始遍历,所以需要边界判断
	if head == nil || head.Next == nil {
		return head
	}
	// 改变指针指向,需要创建两个指针pre,cur指针
	pre, cur := head, head.Next
	for cur != nil {
		if pre.Val == cur.Val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return head
}
