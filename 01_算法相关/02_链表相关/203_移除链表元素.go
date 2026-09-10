package main

/*
解题思路:
因为head节点的Val属性的值可能与val相等;所以这套题的思路毋庸置疑一定是哨兵节点;
因为涉及删除节点,删除节点,必然会导致pre节点和nxt节点的指针改变;所以需要创建两个变量pre,cur进行遍历修改指针的指向
*/
func removeElements(head *ListNode, val int) *ListNode {
	// 创建一个哨兵节点: 因为head节点的Val可能与val相等;
	dummy := &ListNode{Next: head}
	// 开始遍历,遇到Val相等的就行指针的改变;因为涉及指针改变,所以需要定义两个变量pre,cur
	pre, cur := dummy, dummy.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return dummy.Next
}
