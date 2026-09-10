package main

// 725_分隔链表
func splitListToParts(head *ListNode, k int) []*ListNode {
	// 计算出链表的长度
	count := 0
	for cur := head; cur != nil; cur = cur.Next {
		count++
	}
	// 定义结果变量
	ans := make([]*ListNode, k)
	// 求商和求余
	m, n := count/k, count%k
	// 第一个for循环是分成k组
	cur := head
	for i := 0; i < k && cur != nil; i++ {
		ans[i] = cur
		// 计算每组容纳的个数
		size := m
		// 保证每个组之间的相差值不会大于1
		if n > 0 {
			size = size + 1
		}
		// 第二个for循环找到每组的最后一个元素
		for j := 0; j < size-1; j++ {
			cur = cur.Next
		}
		nxt := cur.Next
		cur.Next = nil
		cur = nxt
		if n > 0 {
			n--
		}
	}
	return ans
}
