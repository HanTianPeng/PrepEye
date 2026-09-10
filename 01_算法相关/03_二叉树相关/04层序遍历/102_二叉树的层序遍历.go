package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的层序遍历: dfs算法
func levelOrderTraversal(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	var dfs func(node *TreeNode, level int)
	dfs = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		// 初始化每一层的数组
		if len(ans) == level {
			ans = append(ans, make([]int, 0))
		}
		ans[level] = append(ans[level], node.Val)
		dfs(node.Left, level+1)
		dfs(node.Right, level+1)
	}
	dfs(root, 0)
	return ans
}

// 二叉树的层序遍历: 迭代方法(队列实现)
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	// 创建一个队列
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	// 遍历
	for len(queue) > 0 {
		levelCount := len(queue)
		// 不能直接初始化为make([]int, levelCount); 这样结果就是[0, 0, 1, 2]
		levelAns := make([]int, 0, levelCount)
		for i := 0; i < levelCount; i++ {
			node := queue[i]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			levelAns = append(levelAns, node.Val)
		}
		queue = queue[levelCount:]
		ans = append(ans, levelAns)
	}
	return ans
}

// 二叉树的层序遍历: dfs算法
func levelOrderTraversal2(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	dfs2(root, 0, &ans)
	return ans
}

func dfs2(node *TreeNode, level int, ans *[][]int) {
	if node == nil {
		return
	}
	// 初始化每一层的数组
	if len(*ans) == level {
		*ans = append(*ans, make([]int, 0))
	}
	(*ans)[level] = append((*ans)[level], node.Val)
	dfs2(node.Left, level+1, ans)
	dfs2(node.Right, level+1, ans)
}
