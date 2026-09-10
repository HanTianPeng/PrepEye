package main

/*
1. 最深叶节点只有一个;那么它的最近公共祖先就是它自己
2. 最深叶节点多个;那么它的最近公共祖先到达最深叶节点的高度相同
=> 最深叶节点的最近公共祖先到最深叶节点的高度相同
*/

// 最深叶节点的最近公共祖先: 二叉树的最大深度(带返回值-返回节点深度)
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	// 边界处理
	if root == nil {
		return root
	}
	// 创建一个二叉树的最大深度递归函数,函数有返回值
	maxDepth := 0
	var ans *TreeNode
	var dfs func(node *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		maxDepth = max(depth, maxDepth)
		if node == nil {
			return depth
		}
		leftDepth := dfs(node.Left, depth+1)
		rightDepth := dfs(node.Right, depth+1)

		if leftDepth == rightDepth && leftDepth == maxDepth {
			ans = node
		}
		return max(leftDepth, rightDepth)
	}
	dfs(root, 0)
	return ans
}

func lcaDeepestLeaves2(root *TreeNode) *TreeNode {
	// 边界处理
	if root == nil {
		return root
	}
	// 写一个二叉树的最大深度算法
	var ans *TreeNode
	maxDepth := 0
	var dfs func(node *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		// 递归函数的边界处理
		if node == nil {
			maxDepth = max(depth, maxDepth)
			return depth
		}
		leftDepth := dfs(node.Left, depth+1)
		rightDepth := dfs(node.Right, depth+1)
		if leftDepth == maxDepth && rightDepth == maxDepth {
			ans = node
		}
		return max(leftDepth, rightDepth)
	}
	dfs(root, 0)
	return ans
}
