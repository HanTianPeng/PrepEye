package main

import "slices"

// 从前序与中序遍历序列构建二叉树
func buildTreeV1(preorder []int, inorder []int) *TreeNode {
	// 边界处理
	if len(preorder) == 0 {
		return nil
	}

	// 前序遍历: 根->左->右; 根据前序数组获取根节点
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}

	// 边界处理
	if len(preorder) == 1 {
		return root
	}
	// 中序遍历: 左->根->右; 根据根节点计算中序数组的根节点的位置
	inorderIndex := slices.Index(inorder, rootVal)

	// 根据中序数组的根节点位置对中序数组进行左右子树进行切割
	inorderLeft := inorder[:inorderIndex]
	inorderRight := inorder[inorderIndex+1:]

	// 计算左子树的长度
	lenLeft := len(inorderLeft)

	// 根据左子树的长度对前序数组进行切割
	preorderLeft := preorder[1 : lenLeft+1]
	preorderRight := preorder[lenLeft+1:]

	// 开始构建二叉树
	root.Left = buildTree(preorderLeft, inorderLeft)
	root.Right = buildTree(preorderRight, inorderRight)
	return root
}
