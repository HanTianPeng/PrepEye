package main

// 合并二叉树: 递归方法(前序遍历)
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 1.递归函数的终止条件
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 2.递归函数的逻辑处理
	root := &TreeNode{Val: root1.Val + root2.Val}
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}

// 合并二叉树: 迭代方法(前序遍历-更好理解)
func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 边界处理
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	// 创建一个栈
	stack := []*TreeNode{root2, root1}
	// 前序遍历(栈)
	for len(stack) > 0 {
		node1 := stack[len(stack)-1]
		node2 := stack[len(stack)-2]
		// 栈缩容
		stack = stack[:len(stack)-2]
		// 将node1作为返回节点
		node1.Val += node2.Val
		// node1,node2的右子树都不为空
		if node1.Right != nil && node2.Right != nil {
			stack = append(stack, node2.Right, node1.Right)
		}
		// node1,node2的左子树都不为空
		if node1.Left != nil && node2.Left != nil {
			stack = append(stack, node2.Left, node1.Left)
		}
		// node1的右子树为空; 不用判断node2的右子树为空的情况,因为这个情况node1.Right还是指向自己
		if node1.Right == nil {
			node1.Right = node2.Right
		}
		// node1的左子树为空; 不用判断node2的左子树为空的情况,因为这个情况node1.Left还是指向自己
		if node1.Left == nil {
			node1.Left = node2.Left
		}
	}
	return root1
}
