package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
二叉树节点的深度: 从根节点到该节点的最长简单路径的节点个数或则边的个数;一般使用前序遍历
二叉树节点的高度: 从该节点到叶子节点的最长简单路径的节点个数或则边的个数;一般使用后序遍历
*/

// 判断是否是平衡二叉树: 递归方法(后序遍历: 左->右->根)
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var balanced func(node *TreeNode) int
	balanced = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftHeight := balanced(node.Left)
		if leftHeight == -1 {
			return -1
		}
		rightHeight := balanced(node.Right)
		if rightHeight == -1 {
			return -1
		}
		// 判断左边与右边的高度差值是否大于1
		if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
			// 返回-1是为了方便不需要继续递归左右节点比较
			return -1
		}
		return max(leftHeight, rightHeight) + 1
	}
	return balanced(root) == -1
}

// 判断是否平衡二叉树: 递归方法 + 二叉树的最大深度
func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var maxDepth func(node *TreeNode) int
	maxDepth = func(root *TreeNode) int {
		// 递归边界
		if root == nil {
			return 0
		}
		return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
	}
	// 左子树的最大深度
	leftMaxHeight := maxDepth(root.Left)
	rightMaxHeight := maxDepth(root.Right)
	if leftMaxHeight-rightMaxHeight > 1 || rightMaxHeight-leftMaxHeight > 1 {
		return false
	}
	return isBalanced2(root.Left) && isBalanced2(root.Right)

}
