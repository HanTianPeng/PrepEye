package main

import "math"

// 验证二叉搜索树: 左节点 < 根节点 < 右节点(陷阱: 不能单纯比较左节点<根节点<右节点;而是左子树都小于根节点,根节点都小于右子树)
func isValidBST(root *TreeNode) bool {
	// 边界处理
	if root == nil {
		return true
	}

	var bst func(node *TreeNode, leftMin, rightMax int) bool
	bst = func(node *TreeNode, leftMin, rightMax int) bool {
		if node == nil {
			return true
		}
		// 需要比较左子树所有的节点的值小于根节点的值;右子树所有节点的值大于根节点的值
		return bst(node.Left, leftMin, node.Val) && bst(node.Right, node.Val, rightMax) && node.Val > leftMin && node.Val < rightMax
	}
	return bst(root, math.MinInt64, math.MaxInt64)
}

// 验证二叉搜索树: 利用二叉搜索树的中序遍历(左<根<右)构建的有序递增数组进行判断
func isValidBST2(root *TreeNode) bool {
	// 边界处理
	if root == nil {
		return true
	}
	// 二叉搜索树的中序遍历的变量定义
	var preNode *TreeNode
	curNode := root
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || curNode != nil {
		if curNode != nil {
			stack = append(stack, curNode)
			curNode = curNode.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if preNode != nil && node.Val <= preNode.Val {
				return false
			}
			preNode = node
			curNode = node.Right
		}
	}
	return true
}
