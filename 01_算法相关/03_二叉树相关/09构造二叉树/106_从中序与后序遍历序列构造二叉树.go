package main

import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	// 边界处理
	if len(inorder) == 0 {
		return nil
	}

	// 从后序数组中找到根节点
	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{Val: rootVal}

	// 边界处理
	if len(inorder) == 1 {
		return root
	}
	// 根据根节点从中序数组中切割出左子树和右子树
	rootInOrderIndex := slices.Index(inorder, rootVal)

	// 开始切割中序数组
	leftInOrder := inorder[:rootInOrderIndex]
	rightInOrder := inorder[rootInOrderIndex+1:]

	// 计算左子树的长度
	leftLen := len(leftInOrder)

	// 根据左子树的长度开始切割后序数组
	leftPostOrder := postorder[:leftLen]
	rightPostOrder := postorder[leftLen : len(postorder)-1]

	// 开始构建二叉树
	root.Left = buildTree(leftInOrder, leftPostOrder)
	root.Right = buildTree(rightInOrder, rightPostOrder)
	return root
}

func buildTree2(inorder []int, postorder []int) *TreeNode {
	inorderIndexMap := make(map[int]int, 0)
	for i, v := range inorder {
		inorderIndexMap[v] = i
	}

	var rebuild func(inorderLeftIndex, postorderLeftIndex, postorderRightIndex int) *TreeNode
	rebuild = func(inorderLeftIndex, postorderLeftIndex, postorderRightIndex int) *TreeNode {
		if postorderLeftIndex == postorderRightIndex {
			return nil
		}
		leftSize := inorderIndexMap[postorder[postorderRightIndex-1]] - inorderLeftIndex
		left := rebuild(inorderLeftIndex, postorderLeftIndex, postorderLeftIndex+leftSize) // 左闭右开
		right := rebuild(inorderLeftIndex+leftSize+1, postorderLeftIndex+leftSize, postorderRightIndex-1)
		return &TreeNode{
			Val:   postorder[postorderRightIndex-1],
			Left:  left,
			Right: right,
		}
	}
	return rebuild(0, 0, len(inorder))
}
