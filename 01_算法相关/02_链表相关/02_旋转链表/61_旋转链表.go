package main

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
示例 1：
输入：head = [1,2,3,4,5], k = 2
输出：[4,5,1,2,3]
示例 2：
输入：head = [0,1,2], k = 4
输出：[2,0,1]
提示：
链表中节点的数目在范围 [0, 500] 内
-100 <= Node.val <= 100
0 <= k <= 2 * 109
解题思路:

	成环并断链
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

// 旋转链表
func rotateRight(head *ListNode, k int) *ListNode {
	// 1.边界判断
	if head == nil || k == 0 {
		return head
	}
	// 2.计算链表的长度并找到tail节点
	tail, count := head, 1
	for tail.Next != nil {
		tail = tail.Next
		count++
	}
	// 3.计算k的余数
	k = k % count
	if k == 0 {
		return head
	}
	// 4.将tail与head进行成环
	tail.Next = head
	// 5.找到新的尾节点的前一个节点,新的尾节点是count-k
	target := head
	step := count - k - 1
	for step > 0 {
		target = target.Next
		step--
	}
	// 6.断链
	nxt := target.Next
	target.Next = nil
	return nxt

}
