package main

/*
解题思路: 二叉树的层序遍历
*/
type Node struct {
	Val      int
	Children []*Node
}

func levelOrderN(root *Node) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelCount := len(queue)
		levelAns := make([]int, 0, levelCount)
		for i := 0; i < levelCount; i++ {
			node := queue[i]
			levelAns = append(levelAns, node.Val)
			queue = append(queue, node.Children...)
		}
		queue = queue[levelCount:]
		ans = append(ans, levelAns)
	}
	return ans
}
