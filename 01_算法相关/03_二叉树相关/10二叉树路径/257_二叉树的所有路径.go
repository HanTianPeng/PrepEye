package main

import "strconv"

/*
解题思路: 二叉树的前序遍历
*/
// 二叉树的所有路径: 递归非数组指针版本
func binaryTreePaths(root *TreeNode) []string {
	ans := make([]string, 0)
	res := travel(root, "", ans)
	return res
}

func travel(node *TreeNode, s string, res []string) []string {
	if node.Left == nil && node.Right == nil {
		res = append(res, s+strconv.Itoa(node.Val))
		return res
	}
	// 这里背后隐藏了回溯思想
	s = s + strconv.Itoa(node.Val) + "->"
	if node.Left != nil {
		res = travel(node.Left, s, res)
	}
	if node.Right != nil {
		res = travel(node.Right, s, res)
	}
	return res
}

// 二叉树的所有路径: 递归数组指针版本
func binaryTreePathsV2(root *TreeNode) []string {
	ans := make([]string, 0)
	travelV2(root, "", &ans)
	return ans
}
func travelV2(node *TreeNode, s string, res *[]string) {
	// 叶子节点的判断条件: 左右节点都为空
	if node.Left == nil && node.Right == nil {
		*res = append(*res, s+strconv.Itoa(node.Val))
		return
	}
	// 拼接root节点
	s = s + strconv.Itoa(node.Val) + "->"
	// 继续递归左右节点
	if node.Left != nil {
		travelV2(node.Left, s, res)
	}
	if node.Right != nil {
		travelV2(node.Right, s, res)
	}
}

// 二叉树的所有路径: 迭代栈(前序遍历: 根->左->右)
func binaryTreePathsV3(root *TreeNode) []string {
	ans := make([]string, 0)
	if root == nil {
		return ans
	}
	stack := []*TreeNode{root}
	pathList := []string{""}
	for len(stack) > 0 {
		l := len(stack)
		node := stack[l-1]
		stack = stack[:l-1]
		path := pathList[l-1]
		pathList = pathList[:l-1]
		if node.Left == nil && node.Right == nil {
			ans = append(ans, path+strconv.Itoa(node.Val))
			continue
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			pathList = append(pathList, path+strconv.Itoa(node.Val)+"->")
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			pathList = append(pathList, path+strconv.Itoa(node.Val)+"->")
		}
	}
	return ans
}
