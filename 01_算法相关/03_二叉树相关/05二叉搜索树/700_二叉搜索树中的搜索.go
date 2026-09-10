package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
二叉搜索树:
如果它的左子树不为空,则左子树上所有节点的值均小于它的根节点的值
如果它的右子树不为空,则右子树上所有节点的值均大于它的根节点的值
它的左子树和右子树均为二叉搜索树
*/

// 二叉搜索树中的搜索: 递归方法(利用了二叉搜索树的左节点<根节点<右节点的特性;进行二分查找)
func searchBST(root *TreeNode, val int) *TreeNode {
	// 递归边界处理
	if root == nil {
		// 找不到
		return nil
	}
	// 根节点大,则可能存在左子树上
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	// 根节点小,则可能存在右子树上
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return root
}

// 二叉搜索树中的搜索: 迭代方法(利用了二叉搜索树的左节点<根节点<右节点的特性;进行二分查找)
func searchBST2(root *TreeNode, val int) *TreeNode {
	// 边界处理
	if root == nil {
		return nil
	}
	// 创建一个队列
	queue := []*TreeNode{root}
	// 遍历(队列)
	for len(queue) > 0 {
		node := queue[0]
		// 队列缩容
		queue = queue[1:]
		if node.Val > val && node.Left != nil {
			queue = append(queue, node.Left)
		} else if node.Val < val && node.Right != nil {
			queue = append(queue, node.Right)
		} else if node.Val == val {
			return node
		}
	}
	return nil
}
