package main

/*
解题思路: 二叉树的程序遍历
*/

type Node2 struct {
	Val   int
	Left  *Node2
	Right *Node2
	Next  *Node2
}

func connect(root *Node2) *Node2 {
	if root == nil {
		return root
	}
	// 创建一个队列
	queue := []*Node2{root}
	// 队列遍历
	for len(queue) > 0 {
		levelCount := len(queue)
		for i := 0; i < levelCount; i++ {
			var nxt *Node2
			if i < levelCount-1 {
				nxt = queue[i+1]
			}
			node := queue[i]
			node.Next = nxt
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
