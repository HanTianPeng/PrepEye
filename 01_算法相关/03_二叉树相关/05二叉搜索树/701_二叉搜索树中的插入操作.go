package main

// 二叉搜索树中的插入操作
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	// 二叉搜索树的中序遍历: 左->根->右
	if root == nil {
		return &TreeNode{Val: val}
	}
	var preNode *TreeNode
	curNode := root
	for curNode != nil {
		preNode = curNode
		if curNode.Val > val {
			curNode = curNode.Left
		} else {
			curNode = curNode.Right
		}
	}
	if preNode.Val > val {
		preNode.Left = &TreeNode{Val: val}
	} else {
		preNode.Right = &TreeNode{Val: val}
	}
	return root
}
