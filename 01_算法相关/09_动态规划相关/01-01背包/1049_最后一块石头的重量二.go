package main

/*
有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。
每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。
示例 1：
输入：stones = [2,7,4,1,8,1]
输出：1
解释：
组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
示例 2：
输入：stones = [31,26,33,21,40]
输出：5
提示：
1 <= stones.length <= 30
1 <= stones[i] <= 100
解题思路:
典型的01背包问题:
背包的重量就是：sum/2
商品: stones
用一个背包的重量为sum/2来装石头,装的最大值就是sum - dp[j] - dp[j]的最小值
*/
func lastStoneWeightII(stones []int) int {
	// 计算石头的重量
	sum := 0
	for _, num := range stones {
		sum += num
	}
	// 计算背包的重量
	packageWeight := sum / 2
	// 定义动态规划一维数组
	dp := make([]int, packageWeight+1)
	// 初始化动态规划: dp[0] = 0
	// 动态规划遍历: 先遍历商品,再遍历背包;背包遍历为逆序遍历
	for _, num := range stones {
		for j := packageWeight; j >= num; j-- {
			dp[j] = max(dp[j], dp[j-num]+num)
		}
	}
	return sum - dp[packageWeight] - dp[packageWeight]
}
