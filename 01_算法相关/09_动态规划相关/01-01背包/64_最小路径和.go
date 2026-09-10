package main

import "fmt"

/*
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。
说明：每次只能向下或者向右移动一步。
{1, 3, 1},
{1, 5, 1},
{4, 2, 1},
输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
解题思路:
该题为0/1背包问题
1. 定义dp[j][i]:
*/
func minPathSum(grid [][]int) int {
	// 定义动态规划二维数组
	goodsLen := len(grid[0])
	packageLen := len(grid)
	dp := make([][]int, packageLen)
	for j := range dp {
		dp[j] = make([]int, goodsLen)
		if j == 0 {
			dp[0][0] = grid[0][0]
		} else {
			dp[j][0] = dp[j-1][0] + grid[j][0]
		}
	}
	for i := 1; i < goodsLen; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	fmt.Println(dp) // dp[0][0] = 1 dp[0][1] = 3 dp[0][2] = 1 dp[1][0] = 1 dp[2][0] = 4
	// 初始化
	// 动态规划方程:
	for i := 1; i < goodsLen; i++ {
		for j := 1; j < packageLen; j++ {
			dp[j][i] = min(dp[j-1][i], dp[j][i-1]) + grid[j][i]
		}
	}
	fmt.Println(dp)
	return dp[packageLen-1][goodsLen-1]
}

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	res := minPathSum(grid)
	fmt.Println(res)

}
