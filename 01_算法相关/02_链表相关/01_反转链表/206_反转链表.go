package main

/*
1 -> 2 -> 3
pre=null, cur=1; null <- 1
pre=1, cur=2; null <- 1 <- 2
pre=2, cur=3; null <- 1 <- 2 <- 3
pre=3, cur=null; 返回pre,得到答案 null <- 1 <- 2 <- 3
解题思路
1. 反转链表本质就是将指针的指向发生改变; 前节点和当前节点的指针就行变更
2. 创建一个临时变量缓存下一个节点
3. 进行当前节点和前节点的指针变更
4. 最终返回的是pre节点,因为pre节点成为了整个链表的head节点
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表: 指针改变 + preNode空节点
func reverseList(head *ListNode) *ListNode {
	var preNode *ListNode
	curNode := head
	// 从当前节点开始遍历
	for curNode != nil {
		// 创建一个临时变量;用来缓存下一个遍历的节点
		nxt := curNode.Next
		// 将当前节点的next指针反向指到前一个节点
		curNode.Next = preNode
		// 将前一个节点变更为当前节点
		preNode = curNode
		// 将当前节点变成为下一个节点
		curNode = nxt
	}
	return preNode
}
