package main

/*
小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 node 。
除了 node 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。
给定二叉树的 node 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。
示例 1:
输入: node = [3,2,3,null,3,null,1]
输出: 7
解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
示例 2:
输入: node = [3,4,5,1,3,null,1]
输出: 9
解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9
提示：
树的节点数在 [1, 104] 范围内
0 <= Node.val <= 104
解题思路:
采用二叉树的后序遍历来实现
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 打家劫舍三: 二叉树的后序遍历 + map
func robTree(root *TreeNode) int {
	m := make(map[*TreeNode]int)
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if node.Left == nil && node.Right == nil {
			return node.Val
		}
		// 如果有记录，先查直接返回
		if v, ok := m[node]; ok {
			return v
		}
		// 打劫当前节点
		cur0 := node.Val
		if node.Left != nil {
			cur0 += dfs(node.Left.Left) + dfs(node.Left.Right)
		}
		if node.Right != nil {
			cur0 += dfs(node.Right.Left) + dfs(node.Right.Right)
		}
		// 不打劫当前节点
		cur1 := dfs(node.Left) + dfs(node.Right)
		m[node] = max(cur0, cur1)
		return m[node]
	}
	return dfs(root)
}
