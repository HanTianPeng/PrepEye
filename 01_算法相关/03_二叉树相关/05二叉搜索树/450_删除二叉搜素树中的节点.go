package main

// 删除二叉搜索树中的节点
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 没有找到删除的节点,返回nil
	if root == nil {
		return root
	}
	// 找到了删除的结果
	if root.Val == key {
		// 该节点的左右节点都为nil
		if root.Left == nil && root.Right == nil {
			return nil
		}
		// 该节点的左子树为nil
		if root.Left == nil {
			return root.Right
		}
		// 该节点的右子树为nil
		if root.Right == nil {
			return root.Left
		}
		// 该节点的左右子树都不为nil;找到当前节点的右子树的最小节点(左边)
		cur := root.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		// 当前节点的右子树的最小节点也是比root.Left要大;所以当前左边
		cur.Left = root.Left
		// 将当前节点设置为右子树
		root = root.Right
		// 下面代码可以不需要;因为val不重复
		//root.Right = deleteNode(root.Right, key)
		return root
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	}
	if root.Val < key {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
