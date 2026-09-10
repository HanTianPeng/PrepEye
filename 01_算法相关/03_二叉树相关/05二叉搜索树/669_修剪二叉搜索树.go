package main

// 修剪二叉搜索树: 递归方法
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	// 边界处理
	if root == nil {
		return root
	}
	// 区间返回全部在左侧,保留左侧区间,右侧全部裁剪
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}
	// 区间返回全部在右侧,保留右侧区间,左侧全部裁剪
	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	// 符合范围
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}

// 修剪二叉搜索树: 迭代方法
func trimBST2(root *TreeNode, low int, high int) *TreeNode {
	// 边界处理
	if root == nil {
		return root
	}
	// 处理root,让root移动到区间[low,high]范围内,左闭右闭
	for root != nil && (root.Val > high || root.Val < low) {
		if root.Val < low {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	cur := root
	for cur != nil {
		// 裁剪左边比low还小的,放置到left指针
		for cur.Left != nil && cur.Left.Val < low {
			cur.Left = cur.Left.Right
		}
		cur = cur.Left
	}

	cur = root
	for cur != nil {
		for cur.Right != nil && cur.Right.Val > high {
			cur.Right = cur.Right.Left
		}
		cur = cur.Right
	}
	return root
}
