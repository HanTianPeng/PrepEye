package main

/*
解题思路:
输入: strs = ["10", "0001", "111001", "1", "0"], m = 5, n = 3 输出: 4  最大子集是 {"10","0001","1","0"}
输入: strs = ["10", "0", "1"], m = 1, n = 1 输出: 2  最大的子集是 {"0", "1"}
商品只有一个,选/不选问题: 0/1背包
1.确定dp数组的含义以及下标的含义
	dp[i][j]:最多有i个0,j个1构成的最大子集的大小为dp[i][j]
2.确定递推公式
	dp[i][j] = max(dp[i][j], dp[i-num0][j-num1] + 1)
3.初始化:
	当背包为0的时候,即m,j都为0时候,dp[0][0] = 0
4.确定遍历顺序:
	滚动数组: 背包需要逆序遍历
5.举例推举dp数组
*/

func findMaxForm(strs []string, m int, n int) int {
	// 创建二维度背包的第一个维度m个0
	dp := make([][]int, m+1)
	// 创建二维度背包的第二个维度n个1
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// 初始化: dp[0][0] = 0
	// 动态转换方程式-先遍历商品: dp[i][j] = max(dp[i][j], dp[i - num0][j-num1] + 1)
	for i := 0; i < len(strs); i++ {
		// 计算num0,num1的个数
		num0, num1 := 0, 0
		for _, char := range strs[i] {
			if char == '1' {
				num1++
			} else if char == '0' {
				num0++
			}
		}
		// 遍历背包-倒序遍历
		for j := m; j >= num0; j-- {
			for k := n; k >= num1; k-- {
				dp[j][k] = max(dp[j][k], dp[j-num0][k-num1]+1)
			}
		}
	}
	return dp[m][n]
}
