package main

/*
解题思路: 二叉树的层序遍历
*/
func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	// 边界处理
	if root == nil {
		return ans
	}
	// 创建一个队列
	queue := []*TreeNode{root}
	// 队列的遍历
	for len(queue) > 0 {
		levelCount := len(queue)
		levelAns := make([]int, 0, levelCount)
		for i := 0; i < levelCount; i++ {
			node := queue[i]
			levelAns = append(levelAns, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[levelCount:]
		ans = append([][]int{levelAns}, ans...)
	}
	return ans
}
