package main

/*
解题思路:
1 -> 2 -> 3 -> 4 -> 5
left: 2, right: 4
反转后的结果:
1- > 4 -> 3 -> 2 -> 5
思路如下:
1. 为了避免从head节点开始反转，所以创建一个哨兵节点
2. 先找到2这个节点的pre节点
3. 开始将2 -> 3 -> 4进行翻转链表
4. 第二步的2节点后面要挂5
5. 第二步的2节点的pre节点后面要挂4
*/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}

	// 先找到left这个节点的pre节点
	p0 := dummy
	for i := 0; i < left-1; i++ {
		p0 = p0.Next
	}
	// 反转2 -> 3 -> 4节点; 这个时候cur就是节点5; pre就是节点4
	var pre *ListNode
	cur := p0.Next
	for i := 0; i < right-left+1; i++ {
		nxt := cur.Next
		pre.Next = cur
		pre = cur
		cur = nxt
	}
	// 与上半段进行拼接 2<p0.Next> -> 5    1<p0> -> 4
	p0.Next.Next = cur
	p0.Next = pre
	return dummy.Next
}

// 反转链表二: 哨兵节点 + 反转链表基础 + 指针改变
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	// 哨兵节点: 就是中间的那段需要翻转链表;翻转链表的其中一个变量为pre;如果是从head-1开始翻转,则无法获取pre变量
	dummy := &ListNode{Next: head} // 默认链表是 1->2->3->4->5 增加一个哨兵节点 0->1->2->3->4->5
	pre := dummy
	cur := pre.Next
	// 计算left位置的pre节点和cur节点: 如果left=2,right=4; cur=2, pre=1
	for i := 0; i < left-1; i++ {
		pre = cur
		cur = cur.Next
	}

	// 创建链表翻转的变量
	preNode := pre
	curNode := preNode.Next

	// 找到left-right之间的链表进行翻转: 将 2->3->4 反转成 4->3->2 curNode=5, preNode=4
	for i := 0; i < right-left+1; i++ {
		nxtNode := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = nxtNode
	}

	// 改变前半段指针 1->4;
	pre.Next = preNode
	// 改变后半段指针 2->5
	cur.Next = curNode
	return dummy.Next
}
