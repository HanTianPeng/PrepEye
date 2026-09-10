package main

func mergeKLists(lists []*ListNode) *ListNode {
	// 边界判断: 这个为递归的边界
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	// 二分法: 分而治之
	mid := len(lists) / 2
	left := lists[:mid]
	right := lists[mid:]
	l := mergeKLists(left)
	r := mergeKLists(right)

	// 合并
	return mergeTwoSortedLists(l, r)
}
