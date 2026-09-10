package main

/*
       A
   B       C
D    E   F    G
后序遍历: 左->右->根  D->E->B->F->G->C->A
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的后序遍历: 递归方法(左->右->根)
func postorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, postorderTraversal(root.Left)...)
	ans = append(ans, postorderTraversal(root.Right)...)
	ans = append(ans, root.Val)
	return ans
}

// 二叉树的后序遍历: 迭代方法(左->右->根)
func postorderTraversal2(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	// 创建一个栈
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		// 从栈的顶部弹出一个元素
		node := stack[len(stack)-1]
		ans = append([]int{node.Val}, ans...)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return ans
}
