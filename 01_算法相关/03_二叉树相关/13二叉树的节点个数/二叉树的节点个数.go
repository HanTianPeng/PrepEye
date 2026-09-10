package main

// 二叉树的节点个数: 递归方法: 时间复杂度O(n) 空间复杂度O(logn)
func countNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes1(root.Left) + countNodes1(root.Right) + 1
}
