package main

/*
解题思路: 二叉树的程序遍历
写法可以和116一样;这里提供了一种层级降低内存的写法
*/

func connect2(root *Node2) *Node2 {
	if root == nil {
		return root
	}
	// 创建一个队列
	queue := []*Node2{root}
	// 队列遍历
	for len(queue) > 0 {
		// 前拷贝
		levelQueue := queue
		queue = nil
		for i, node := range levelQueue {
			if i > 0 {
				levelQueue[i-1].Next = node
			}
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
