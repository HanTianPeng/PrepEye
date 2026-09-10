package main

/*
解题思路: 二叉树的层序遍历 + 左右节点都为nil
*/
func minDepth(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelCount := len(queue)
		ans++
		isEnd := false
		for i := 0; i < levelCount; i++ {
			node := queue[i]
			if node.Left == nil && node.Right == nil {
				isEnd = true
				break
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[levelCount:]
		if isEnd {
			break
		}
	}
	return ans
}
