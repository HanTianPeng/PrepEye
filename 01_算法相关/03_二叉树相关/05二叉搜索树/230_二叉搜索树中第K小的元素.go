package main

// 二叉搜索树中第K小的元素: 二叉树的中序遍历迭代方法(本质是因为二叉搜索树的中序遍历就是一个递增的数组)
func kthSmallest(root *TreeNode, k int) int {
	// 边界处理
	if root == nil {
		return 0
	}
	// 二叉树的中序遍历的变量定义
	cur, stack := root, make([]*TreeNode, 0)
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k--
			if k == 0 {
				return node.Val
			}
			cur = node.Right
		}
	}
	return 0
}
