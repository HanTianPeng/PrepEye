package main

/*
奇偶链表
解题思路:
1 -> 2 -> 3 -> 4 -> 5
最终结果
1 -> 3 -> 5 -> 2 -> 4
核心思想前后指针
*/
func oddEvenList(head *ListNode) *ListNode {
	// 边界判断: 如果只有一个节点,也就不存在奇偶链表的情况
	if head == nil || head.Next == nil {
		return head
	}
	// 创建两个变量odd:奇数; even:偶数;
	odd, even := head, head.Next
	// 创建一个变量存偶数的head节点
	evenHead := even
	// 遍历偶数节点
	for even != nil && even.Next != nil {
		// 奇数1的next指向的是3,而3是偶数的next节点
		odd.Next = even.Next
		// 跳转到下一个奇数3
		odd = odd.Next
		// 偶数2的next指向的是4,而4是下一个奇数3的next节点
		even.Next = odd.Next
		// 跳转到下一个偶数
		even = even.Next
	}
	// 将奇数的尾节点的next指向偶数的head节点
	odd.Next = evenHead
	return head
}
