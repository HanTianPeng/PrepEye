package main

/*
将二叉树以前序遍历的顺序进行展开
*/

// 二叉树展开为链表: 递归方法(前序遍历)
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var preNode *TreeNode
	var dfs func(curNode *TreeNode)
	dfs = func(curNode *TreeNode) {
		// 递归的终止条件
		if curNode == nil {
			return
		}
		if preNode != nil {
			preNode.Right = curNode.Right
			preNode.Left = nil
		}
		preNode = curNode
		dfs(curNode.Left)
		dfs(curNode.Right)
	}
	dfs(root)
}

// 二叉树展开为链表: 迭代方法(前序遍历)
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	var preNode *TreeNode
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if preNode != nil {
			preNode.Right = node
			preNode.Left = nil
		}
		preNode = node
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return
}
