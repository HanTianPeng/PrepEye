package main

import "slices"

// 路径总和二: 递归
func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	var path []int
	// 1-确定递归函数的入参和出参
	var dfs func(node *TreeNode, targetSum int)

	dfs = func(node *TreeNode, targetSum int) {
		// 2-确定递归函数的终止条件
		if node == nil {
			return
		}
		path = append(path, node.Val)
		targetSum -= node.Val
		if node.Left == nil && node.Right == nil && targetSum == 0 {
			ans = append(ans, slices.Clone(path))
		}
		dfs(node.Left, targetSum)
		dfs(node.Right, targetSum)
		path = path[:len(path)-1] // A -> B 处理完成   恢复现场方便 A -> C
	}
	dfs(root, targetSum)
	return ans
}

// 路径总和二: 迭代方法(栈 + slices深拷贝)
func pathSum2(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	// 创建一个栈
	stack := []*TreeNode{root}
	sumStack := []int{targetSum}
	pathList := make([][]int, 0)
	pathList = append(pathList, []int{})
	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		stack = stack[:l-1]
		sum := sumStack[l-1]
		sumStack = sumStack[:l-1]
		pathItem := pathList[l-1]
		pathList = pathList[:l-1]
		if node.Left == nil && node.Right == nil && sum == node.Val {
			pathItem = append(pathItem, node.Val)
			res = append(res, pathItem)
			continue
		}
		if node.Right != nil {
			sumStack = append(sumStack, sum-node.Val)
			stack = append(stack, node.Right)
			i := slices.Clone(pathItem)
			i = append(i, node.Val)
			pathList = append(pathList, i)
		}

		if node.Left != nil {
			sumStack = append(sumStack, sum-node.Val)
			stack = append(stack, node.Left)
			i := slices.Clone(pathItem)
			i = append(i, node.Val)
			pathList = append(pathList, i)
		}
	}
	return res
}
