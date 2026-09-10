package main

// 二叉搜索树中的众数: 二叉树的中序遍历: 左->根->右
func findMode(root *TreeNode) []int {
	ans := make([]int, 0)
	// 边界处理
	if root == nil {
		return ans
	}
	// 二叉搜索树的中序遍历的特点: 有序数组
	var preNode *TreeNode
	// 针对于有序数组的众数统计的变量定义
	count, maxCount := 0, 0
	// 二叉树中序遍历的变量定义
	cur := root
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			// 从栈中获取节点
			node := stack[len(stack)-1]
			// 栈缩容
			stack = stack[:len(stack)-1]
			// 处理逻辑
			if preNode == nil {
				count = 1
			} else if preNode.Val == node.Val {
				count++
			} else {
				count = 1
			}
			// 结果判断
			if count == maxCount {
				ans = append(ans, node.Val)
			}
			if count > maxCount {
				maxCount = count
				ans = []int{node.Val}
			}
			// 更新preNode节点
			preNode = node
			// 右子树的迭代
			cur = node.Right
		}
	}
	return ans
}
