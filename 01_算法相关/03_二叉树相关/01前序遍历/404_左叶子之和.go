package main

// 左叶子节点之和: 递归方法(左叶子节点: 左节点不为nil,左节点的左右节点都为nil)
func sumOfLeftLeaves(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}
	left, right := root.Left, root.Right
	if left != nil && left.Left == nil && left.Right == nil {
		sum += left.Val
	}
	return sum + sumOfLeftLeaves(left) + sumOfLeftLeaves(right)
}

// 左叶子之和: 迭代方法(前序遍历: 栈)
func sumOfLeftLeaves2(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}
	// 创建一个栈
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 左叶子节点处理
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			sum += node.Left.Val
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return sum
}
