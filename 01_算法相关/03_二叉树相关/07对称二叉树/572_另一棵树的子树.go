package main

// 另一棵树的子树: 相同二叉树 + 双层递归
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// 递归的结束条件
	if root == nil {
		return false
	}
	// 根节点和subRoot相同  递归判断sub节点
	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
