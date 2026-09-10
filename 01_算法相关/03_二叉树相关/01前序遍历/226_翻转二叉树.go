package main

/*
解题思路: 二叉树的前序遍历
*/

// 翻转二叉树 递归方法: 更改指针
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 翻转二叉树 迭代方法: 前序遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		temp := node.Left
		node.Left = node.Right
		node.Right = temp
	}
	return root
}
