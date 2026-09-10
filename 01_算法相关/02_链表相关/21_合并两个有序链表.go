package main

/*
解题思路: 递归
谁小就把谁作为当前节点
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 边界判断
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	// 比较大小;谁小就以谁作为当前节点;并且用小的节点的next节点与另外一个节点进行递归比较
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}

/*
解题思路: 迭代 + 哨兵节点
*/
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	// 创建一个哨兵节点
	dummy := &ListNode{}
	// 遍历list1,list2;比较谁小就挂在哨兵节点的下面
	curNode := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curNode.Next = list1
			list1 = list1.Next
		} else {
			curNode.Next = list2
			list2 = list2.Next
		}
		curNode = curNode.Next
	}

	// 判断list1还没有遍历完
	if list1 != nil {
		curNode.Next = list1
	}
	// 判断list2还没有遍历完
	if list2 != nil {
		curNode.Next = list2
	}
	return dummy.Next
}
