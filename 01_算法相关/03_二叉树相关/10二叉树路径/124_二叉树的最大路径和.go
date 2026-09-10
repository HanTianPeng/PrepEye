package main

import "math"

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt
	if root == nil {
		return 0
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		// 计算左边分支的最大值:左边分支如果为负数,则不如不选择
		leftVal := max(0, dfs(node.Left))
		// 计算右边分支的最大值:右边分支如果为负数,则不如不选择
		rightVal := max(0, dfs(node.Right))
		// 路径的最大值: 不需要去单独比较leftVal, rightVal;因为leftVal或rightVal为0时候,证明都没有被选择
		ans = max(ans, leftVal+rightVal+ans)

		// 返回经过root的单边最大分支给当前root的父节点计算使用
		return max(leftVal, rightVal) + node.Val
	}
	dfs(root)
	return ans
}
