package main

/*
		1
	2		3
4		5
二叉树的直径: 4 - 2 - 1 - 3
*/

// 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	var ans int
	// 定义一个递归函数，返回左右两棵树的最大值
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		// 判断叶子节点: -1 + 1 = 0;因为需要结果为边
		if node == nil {
			return -1
		}
		left := dfs(node.Left) + 1
		right := dfs(node.Right) + 1
		ans = max(ans, left+right)

		return max(left, right) // 当前节点的最大边
	}
	return ans
}
