package main

func maxDepth2(root *Node) int {
	ans := 0
	if root == nil {
		return ans
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		levelCount := len(queue)
		for i := 0; i < levelCount; i++ {
			node := queue[i]
			queue = append(queue, node.Children...)
		}
		queue = queue[levelCount:]
		ans++
	}
	return ans
}
