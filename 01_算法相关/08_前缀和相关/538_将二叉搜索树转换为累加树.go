package main

/*
给出二叉 搜索 树的根节点 root，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），
将其转换为一个更大的树，使得原始二叉搜索树中的每个节点值都变为原本值加上原本二叉搜索树中所有比该节点值大的节点值的总和。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将二叉搜索树转化为累加树: 二叉搜索树 + 中序遍历(左 -> 中 -> 右)的逆序
func convertBST(root *TreeNode) *TreeNode {
	// 创建一个前缀和变量
	preSum := 0
	// 创建一个二叉树的中序遍历(逆序)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		// 递归的边界处理
		if node == nil {
			return
		}
		// 右子树遍历
		dfs(node.Right)
		preSum += node.Val
		node.Val = preSum
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func convertBST2(root *TreeNode) *TreeNode {
	// 边界处理
	if root == nil {
		return root
	}
	// 二叉树的中序遍历(逆序)
	cur, stack := root, make([]*TreeNode, 0)
	// 创建前缀和的变量
	preSum := 0
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Right
		} else {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			preSum += node.Val
			node.Val = preSum
			cur = node.Left
		}
	}
	return root
}
