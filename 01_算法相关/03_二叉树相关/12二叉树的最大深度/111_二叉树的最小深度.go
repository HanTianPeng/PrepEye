package main

// 二叉树的最小深度: 递归
func minDepth1(root *TreeNode) int {
	// 边界处理
	if root == nil {
		return 0
	}
	minDepth := 0
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		minDepth = min(depth, minDepth)
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			minDepth = min(minDepth, depth)
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 1)
	return minDepth
}

// 二叉树的最小深度: 递归方法
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth2(root.Right)
	}
	if root.Left != nil && root.Right == nil {
		return 1 + minDepth2(root.Left)
	}
	return 1 + min(minDepth2(root.Left), minDepth2(root.Right))
}
