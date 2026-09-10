package main

/*
解题思路:
由于参数是被删除的节点;
前提条件:

	链表的所有的节点的值互不相等；
	被删除的节点不是链表的最后一个节点;

所以可以通过改变Val属性来完成删除链表的操作
*/
func deleteNode(node *ListNode) {
	// 将当前node值修改成下一个node的值
	node.Val = node.Next.Val
	// 将当前node的next指针指向下一个node节点的下一个节点
	node.Next = node.Next.Next
}
