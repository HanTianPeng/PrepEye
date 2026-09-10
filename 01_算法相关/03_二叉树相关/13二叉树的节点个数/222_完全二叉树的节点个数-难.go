package main

/*
解题思路:
1. 当成一个普通二叉树,用层序或则其他的遍历方法都可以解决
2. 利用完全二叉树的特性解决
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 完全二叉树的节点个数: 利用了完全二叉树的特性: 时间复杂度: O(logn * logn), 空间复杂度O(logn)
func countNodes(root *TreeNode) int {
	// 递归边界处理
	if root == nil {
		return 0
	}
	leftH, rightH := 0, 0
	left, right := root.Left, root.Right
	// 遍历统计左边节点的高度
	for left != nil {
		leftH++
		left = left.Left
	}
	for right != nil {
		rightH++
		right = right.Right
	}
	// 如果左右高度一致,证明是满二叉树,满二叉树的节点数:2^k - 1
	if leftH == rightH {
		return 2<<leftH - 1
	}
	// 如果左右高度不一致,就分别去计算左右的高度: 最终肯定会遇到满二叉树
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
