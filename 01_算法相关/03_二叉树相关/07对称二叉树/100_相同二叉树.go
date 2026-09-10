package main

// 相同的二叉树: 递归方法
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 递归的终止条件
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 相同的二叉树: 迭代方法
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	// 创建一个队列
	queue := []*TreeNode{p, q}
	// 队列遍历
	for len(queue) > 0 {
		pNode := queue[0]
		qNode := queue[1]
		queue = queue[2:]
		// 这里判断后,就不需要进行边界处理
		if pNode == nil && qNode == nil {
			continue
		}
		if pNode == nil || qNode == nil || pNode.Val != qNode.Val {
			return false
		}
		// 因为前面的判断已经判断了pNode,qNode都不可能为nil;所以直接append
		queue = append(queue, pNode.Left, qNode.Left, pNode.Right, qNode.Right)
	}
	return true
}
