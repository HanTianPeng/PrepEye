package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 如果链表A和链表B存在相交节点: a1 + c1 + b1 = b1 + c1 + a1 最后两个指针都指向交点;此时刚好返回value
// 如果链表A和链表B不存在相交节点: a1 + b1 = b1 + a1 最后两个指针都指向null;此时也相等;刚好返回null
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}
