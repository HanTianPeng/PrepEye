package main

type MyLinkedList struct {
	head *ListNode
	size int
}

// Constructor1 初始化单向链表; 默认创建一个哨兵节点,并且哨兵节点不算size
func Constructor1() MyLinkedList {
	// 初始化单向链表的时候,默认创建一个哨兵节点,但是把size设置为0
	return MyLinkedList{&ListNode{}, 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head
	// 因为已经存在了哨兵节点,所以i <= index
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	// 边界处理
	if index > this.size {
		return
	}

	// 针对于index <= 0 都按照0索引添加节点
	index = max(index, 0)
	this.size++
	pre := this.head
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	addNode := &ListNode{}
	addNode.Val = val
	addNode.Next = pre.Next
	pre.Next = addNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	this.size--
	pre := this.head
	for i := 0; i < index; i++ {
		pre = pre.Next
	}
	// 改变指针
	pre.Next = pre.Next.Next
}
