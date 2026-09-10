package main

/*
解题思路: 二叉树的后序遍历
*/

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	for _, node := range root.Children {
		ans = append(ans, postorder(node)...)
	}
	ans = append(ans, root.Val)
	return ans
}

func postorder2(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	// 创建一个栈
	stack := []*Node{root}
	for len(stack) > 0 {
		// 当前节点
		node := stack[len(stack)-1]
		// 栈缩容
		stack = stack[:len(stack)-1]
		ans = append([]int{node.Val}, ans...)
		// 这个循环注意了,和后序遍历不一样
		for _, child := range node.Children {
			stack = append(stack, child)
		}

	}
	return ans
}
