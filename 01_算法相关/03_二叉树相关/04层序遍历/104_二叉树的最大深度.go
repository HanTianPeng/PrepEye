package main

/*
二叉树节点的深度: 从根节点到该节点的最长简单路径的节点个数或则边的个数;一般使用前序遍历
二叉树节点的高度: 从该节点到叶子节点的最长简单路径的节点个数或则边的个数;一般使用后序遍历
解题思路: 二叉树的层序遍历
*/
// 二叉树的最大深度: 迭代方法(层序遍历)
func maxDepth(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
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
		}
		ans++
		queue = queue[levelCount:]
	}
	return ans
}
