package main

/*
	    A
	B       C

D    E   F    G
中序遍历: 左->根->右  D->B->E->A->F->C->G
二叉树的中序遍历: 左 -> 根 -> 右
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的中序遍历: 递归方法(左->根->右)
func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, inorderTraversal(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversal(root.Right)...)
	return ans
}

// 二叉树的中序遍历:左->根->右
func inorderTraversal2(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	stack := make([]*TreeNode, 0)
	// 创建一个当前指针节点
	cur := root
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			// 将当前指针节点指向左节点
			cur = cur.Left
		} else {
			// 从栈的顶部弹出一个节点
			node := stack[len(stack)-1]
			// 将栈的大小进行缩容
			stack = stack[:len(stack)-1]
			ans = append(ans, node.Val)
			// 将当前指针节点指向右节点
			cur = node.Right
		}
	}
	return ans
}
