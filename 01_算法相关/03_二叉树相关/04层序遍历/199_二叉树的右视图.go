package main

/*
解题思路: 二叉树的层序遍历
*/
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	// 创建一个队列
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	// 遍历队列
	for len(queue) > 0 {
		levelCount := len(queue)
		for i := 0; i < levelCount; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if i == levelCount-1 {
				ans = append(ans, node.Val)
			}
		}
		queue = queue[levelCount:]
	}
	return ans
}
