package main

/*
解题思路: 二叉树的层序遍历
*/
func averageOfLevels(root *TreeNode) []float64 {
	ans := make([]float64, 0)
	if root == nil {
		return ans
	}
	// 创建一个队列
	queue := make([]*TreeNode, 0)
	// 将根节点入到队列中
	queue = append(queue, root)
	// 遍历队列
	for len(queue) > 0 {
		levelCount := len(queue)
		var sum float64
		// 每一层都遍历
		for i := 0; i < levelCount; i++ {
			node := queue[i]
			sum += float64(node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, sum/float64(levelCount))
		queue = queue[levelCount:]
	}
	return ans
}
