package main

/*
解题思路: 二叉树的前序遍历
*/

type Node struct {
	Val      int
	Children []*Node
}

// 递归方法
func preorder(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	for _, node := range root.Children {
		ans = append(ans, preorder(node)...)
	}
	return ans
}

// 迭代方法: 栈
func preorder2(root *Node) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return ans
}
