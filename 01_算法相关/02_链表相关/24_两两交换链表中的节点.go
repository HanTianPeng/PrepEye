package main

/*
解题思路:
1 -> 2 -> 3 -> 4
首先想到的就是哨兵节点: 0
第一轮互换位置:

	0作为preNode, 1作为curNode
	0 -> 2
	2 -> 1
	1 -> 3
	最终结果是 0 -> 2 -> 1 -> 3 -> 4

第二轮互换位置:

	1作为preNode, 3作为curNode
	1 -> 4
	4 -> 3
	3 ->
	最终结果是 0 -> 2 -> 1 -> 4 -> 3 ->
*/
func swapPairs(head *ListNode) *ListNode {
	// 创建一个哨兵节点
	dummy := &ListNode{Next: head}

	// 定义0/1两个变量
	node0, node1 := dummy, head

	// 开始遍历: 为了保证2->1; 2要挂1;那么2就不能为nil
	for node1 != nil && node1.Next != nil {
		node2 := node1.Next
		node3 := node2.Next

		// 第一轮开始指针变更
		node0.Next = node2
		node2.Next = node1
		node1.Next = node3

		// 第二轮
		node0 = node1
		node1 = node3
	}
	return dummy.Next
}
