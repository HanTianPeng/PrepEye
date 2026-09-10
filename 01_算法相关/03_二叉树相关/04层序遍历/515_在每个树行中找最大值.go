package main

import "math"

/*
解题思路: 二叉树的层序遍历 + math.MinInt32
*/
func largestValues(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelCount := len(queue)
		levelMaxNum := math.MinInt32
		for i := 0; i < levelCount; i++ {
			node := queue[i]
			levelMaxNum = max(levelMaxNum, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, levelMaxNum)
		queue = queue[levelCount:]
	}
	return ans
}
