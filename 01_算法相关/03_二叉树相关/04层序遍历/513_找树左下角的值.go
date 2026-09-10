package main

// 找树左下角的值: 层序遍历(迭代方法), 一层更新一次ans值
func findBottomLeftValue(root *TreeNode) int {
	// 二叉树的层序遍历: 创建一个队列
	var ans int
	if root == nil {
		return ans
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelCount := len(queue)
		for i := 0; i < levelCount; i++ {
			node := queue[i]
			if i == 0 {
				ans = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		// 将队列进行缩容
		queue = queue[levelCount:]
	}
	return ans
}

// 找树左下角的值: 递归(优先递归左子树) 判断深度最大的叶子节点: 二叉树的最大深度 + 叶子节点的判断
func findBottomLeftValue2(root *TreeNode) int {
	depth, ans := 0, 0
	// 1-确定递归函数的参数和返回值
	var dfs2 func(node *TreeNode, depth int)
	dfs2 = func(node *TreeNode, d int) {
		// 2-确定递归的终止条件
		if node == nil {
			return
		}
		// 3-确定单层递归的逻辑: 因为先遍历左节点,左节点如果有值,则右节点同层不可能有值
		if node.Left == nil && node.Right == nil && depth < d {
			depth = d
			ans = node.Val
		}
		// 左
		dfs2(node.Left, d+1)
		// 右
		dfs2(node.Right, d+1)
	}
	dfs2(root, 1)
	return ans
}
