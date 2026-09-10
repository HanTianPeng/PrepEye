package main

/*
解题思路:
1. 从题目看出新的链表所有的val都可能发生变化;所以这个时候需要用哨兵节点最合适
2. 两个链表从头节点开始都当成个位数;这个时候适合用双指针
3. 两个数相加大于10，这个时候就需要考虑是否进一的问题
4. 新产生的节点的val等于前面节点进一的数字加上当前两个节点的val除以10取余数
5. 最后节点全部遍历完成,还有可能再进一
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 创建一个哨兵节点
	dummy := &ListNode{}
	// 从当前节点开始挂载新的节点
	cur := dummy
	// 定义进一的变量
	m := 0

	// 双指针遍历l1,l2
	for l1 != nil || l2 != nil || m != 0 {
		// 双指针对应位置和进一的变量进行相加
		if l1 != nil {
			m += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			m += l2.Val
			l2 = l2.Next
		}
		// 除以10求余数作为新节点的val
		cur.Next = &ListNode{Val: m % 10}
		// 重新计算进一的变量
		m = m / 10
		cur = cur.Next
	}
	return dummy.Next
}
