package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 路径总和: 递归
func hasPathSum(root *TreeNode, targetSum int) bool {
	// 1-确定递归函数的参数和返回值

	// 2-确定递归函数的终止条件
	if root == nil {
		return false
	}
	// 3-确定递归函数的逻辑: 根是叶子节点: 根遍历
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	// 左节点或则右节点一个满足条件都可以(左->右)
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum)
}

// 路径总和: 迭代方法(栈, 参考二叉树的所有路径)
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 创建一个栈
	stack := []*TreeNode{root}
	sumStack := []int{targetSum}
	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		stack = stack[:l-1]
		sum := sumStack[l-1]
		sumStack = sumStack[:l-1]
		if node.Left == nil && node.Right == nil && sum == node.Val {
			return true
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			sumStack = append(sumStack, sum-node.Val)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			sumStack = append(sumStack, sum-node.Val)
		}
	}
	return false
}
