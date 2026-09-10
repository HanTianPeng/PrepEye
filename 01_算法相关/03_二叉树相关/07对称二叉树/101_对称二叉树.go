package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 看到对称二叉树,就应该想到617_合并二叉树;递归的参数为左子树和右子树;或则一次入栈的是左子树和右子树

// 对称二叉树: 递归方法
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(left, right *TreeNode) bool
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		// 3-单层递归的逻辑: 比较外侧相等且内侧也相等
		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	return dfs(root.Left, root.Right)
}

// 对称二叉树: 迭代方法(栈)
func isSymmetric2(root *TreeNode) bool {
	// 边界处理
	if root == nil {
		return true
	}
	stack := []*TreeNode{root.Right, root.Left}
	for len(stack) > 0 {
		left := stack[len(stack)-1]
		right := stack[len(stack)-2]
		stack = stack[:len(stack)-2]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		stack = append(stack, right.Left, left.Right, right.Right, left.Left)
	}
	return true
}

// 对称二叉树: 迭代方法(队列,栈都可以,但是通常写队列大家更容易理解)
func isSymmetric3(root *TreeNode) bool {
	// 边界处理
	if root == nil {
		return true
	}
	// 创建一个队列: 但是本道题不是层序遍历;只是为了比较左右两棵树方便
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
