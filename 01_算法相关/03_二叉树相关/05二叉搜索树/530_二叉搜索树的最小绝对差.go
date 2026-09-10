package main

import "math"

/*
二叉搜索树: 等同于有序数组
       6
   4         9
3     5   8     10
二叉搜索树的中序遍历: 左->根->右(3 -> 4 -> 5 -> 6 -> 8 -> 9 -> 10)
解题思路: 二叉搜索树的中序遍历有一个特别的写法
preNode变量,preNode一定小于curNode
*/

// 二叉搜索树的最小绝对差
func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt64
	preVal := math.MinInt64
	// 1-确定递归函数的入参和出参
	var bst func(node *TreeNode)
	bst = func(node *TreeNode) {
		if node == nil {
			return
		}
		bst(node.Left)
		ans = min(node.Val-preVal, ans)
		preVal = node.Val
		bst(node.Right)
	}
	bst(root)
	return ans
}

// 二叉搜索树的最小绝对差: 中序遍历(迭代方法)
func getMinimumDifference2(root *TreeNode) int {
	// 边界处理
	if root == nil {
		return 0
	}
	ans := math.MaxInt64
	var preNode *TreeNode
	// 中序遍历需要创建两个变量
	cur := root
	stack := make([]*TreeNode, 0)
	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			node := stack[len(stack)-1]
			// 栈缩容
			stack = stack[:len(stack)-1]
			// 比较差值
			if preNode != nil {
				ans = min(node.Val-preNode.Val, ans)
			}
			preNode = node
			cur = node.Right
		}
	}
	return ans
}
