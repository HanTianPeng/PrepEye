package main

type Node struct {
	Val      int
	Children []*Node
}

// N叉树的最大深度 递归方法(前序遍历: 根->左->右) + children长度为0的时候边界处理
func maxDepthN1(root *Node) int {
	maxDepth := 0
	var dfs func(node *Node, depth int)
	dfs = func(node *Node, depth int) {
		maxDepth = max(maxDepth, depth)
		if node == nil {
			return
		}
		// 边界处理
		if len(node.Children) == 0 {
			dfs(nil, depth+1)
		}
		for _, childNode := range node.Children {
			dfs(childNode, depth+1)
		}
	}
	dfs(root, 0)
	return maxDepth
}
