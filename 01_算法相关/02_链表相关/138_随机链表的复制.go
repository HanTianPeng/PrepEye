package main

/*
解题思路:

	原链表: a -> b -> c
	第一步原链表基础进行复制,但复制的节点不复制其random_index的字段: a -> A -> b -> B -> c -> C
	第二步复制其random_index字段，因为原节点的next节点都是复制节点;所以只需要改变指针即可
	第三步其实就是为奇偶链表，将链表进行断链操作
*/
func copyRandomList(head *Node) *Node {
	// 复制每一个节点，并且把新的节点直接插入到原链表中
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
		cur = temp
	}

	// 遍历交错链表的原链表的节点
	curNode := head
	for curNode != nil {
		if curNode.Random != nil {
			curNode.Next.Random = curNode.Random.Next
		}
		curNode = curNode.Next.Next
	}

	// 分离交错的链表,创建一个哨兵节点
	dummy := &Node{}
	copyNode := dummy
	for head != nil {
		cloneNode := head.Next
		copyNode.Next = cloneNode
		copyNode = copyNode.Next
		head.Next = cloneNode.Next
		head = cloneNode.Next
	}
	return dummy.Next
}
