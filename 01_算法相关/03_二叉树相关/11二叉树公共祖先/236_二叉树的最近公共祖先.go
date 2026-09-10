package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的最近公共祖先: 后序遍历
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 边界处理: 没有找到是nil
	if root == nil {
		return nil
	}
	// 找到了p或者q,直接返回,这样后续判断直接用这个节点是否为nil进行判断
	if root == p || root == q {
		return root
	}
	// 进行左->右->根后序遍历
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 根节点进行递归的逻辑处理
	// 1. 左右子树都不为null,证明p,q分别在左子树或右子树;公共祖先为root
	if left != nil && right != nil {
		return root
	}
	// 2. 左子树不为null,证明p,q都在左子树;返回最后一个找到的节点
	if left != nil {
		return left
	}
	// 3. 右子树不为null,证明p,q都在右子树;返回最后一个找到的节点
	if right != nil {
		return right
	}
	// 4. 都没有找到
	return nil
}
