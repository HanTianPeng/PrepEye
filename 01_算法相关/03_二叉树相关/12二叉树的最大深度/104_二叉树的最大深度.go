package main

/*
二叉树节点的深度(从上往下): 从根节点到该节点的最长简单路径的节点个数或则边的个数;一般使用前序遍历
二叉树节点的高度(从下往上): 从该节点到叶子节点的最长简单路径的节点个数或则边的个数;一般使用后序遍历
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的最大深度: 递归方法(前序遍历: 根->左->右)
func maxDepth1(root *TreeNode) int {
	// 1-确定递归函数的参数和返回值
	maxDepth := 0
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		// 3-确定递归函数的逻辑
		maxDepth = max(maxDepth, depth)
		// 2-确定递归函数的终止条件
		if node == nil {
			return
		}
		// 左
		dfs(node.Left, depth+1)
		// 右
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return maxDepth
}

// 二叉树的最大深度 递归方法
func maxDepth2(root *TreeNode) int {
	maxDepth := 0
	var dfs func(node *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			return depth - 1
		}
		leftDepth := dfs(node.Left, depth+1)
		rightDepth := dfs(node.Right, depth+1)
		// 这里进行maxDepth是因为需要在depth-1之后进行
		maxDepth = max(maxDepth, leftDepth, rightDepth)
		return max(leftDepth, rightDepth)
	}
	dfs(root, 1)
	return maxDepth
}

// 二叉树的最大深度: 递归方法(后序遍历)
func maxDepth3(root *TreeNode) int {
	// 递归边界
	if root == nil {
		return 0
	}
	return max(maxDepth3(root.Left), maxDepth3(root.Right)) + 1
}

// ----------------不成熟期的写法------------------
// 二叉树的最大深度: 递归方法(根->左->右)
func maxDepth4(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}
	// 1-确定递归函数的参数和返回值
	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		result = max(result, depth)
		// 根节点
		if node.Left == nil && node.Right == nil {
			return
		}
		// 左节点
		if node.Left != nil {
			dfs(node.Left, depth+1)
		}
		// 右节点
		if node.Right != nil {
			dfs(node.Right, depth+1)
		}
		return
	}
	dfs(root, 1)
	return result
}

// 二叉树的最大深度: 递归方法(前序遍历)
func maxDepth5(root *TreeNode) int {
	result := 0
	if root == nil {
		return result
	}
	return getDepth(root, 1, 1)
}

func getDepth(node *TreeNode, depth, result int) int {
	// 中间节点
	if result < depth {
		result = depth
	}
	if node.Left == nil && node.Right == nil {
		return result
	}
	// 左边节点
	if node.Left != nil {
		result = getDepth(node.Left, depth+1, result)
	}
	// 右边节点
	if node.Right != nil {
		result = getDepth(node.Right, depth+1, result)
	}
	return result
}
