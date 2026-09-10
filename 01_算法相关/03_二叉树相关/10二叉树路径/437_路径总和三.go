package main

/*
用坐标画图解释
root---------node1------node2
从root到node1的路径和的个数为n
从root到node2的路径和的个数为m个
那么从node1到node2的路径和的个数为m-n个

假设从根节点到终点node的路径和为s
那么起点到这个终点node的路径和为target的个数: m[s - target]
*/

// 路径总和三: 前缀和 + 递归
func pathSumV1(root *TreeNode, targetSum int) int {
	ans := 0
	// 统计前缀和(从根节点到node节点的路径和出现的次数)
	m := map[int]int{0: 1}
	var dfs func(node *TreeNode, sum int)
	dfs = func(node *TreeNode, sum int) {
		// 1-递归函数的终止条件
		if node == nil {
			return
		}
		sum += node.Val
		// 3-递归函数的逻辑处理
		ans += m[sum-targetSum]
		m[sum]++
		dfs(node.Left, sum)
		dfs(node.Right, sum)
		m[sum]--
	}
	dfs(root, 0)
	return ans
}
