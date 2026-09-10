package main

import "math"

/*
给你一个非负整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
示例 1：
输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：
输入：nums = [1], target = 1
输出：1
提示：
1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000
解题思路:
假设加法个数的总和为x，则剩余全部是减法对应的总和: sum - x
所以得出公式: x - (sum-x) = target => x = (sum + target) / 2
*/

func findTargetSumWays(nums []int, target int) int {
	// 计算数字的最大和
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 边界处理
	if int(math.Abs(float64(target))) > sum {
		return 0
	}
	// 这块没有搞懂
	if (sum+target)%2 != 0 {
		return 0
	}
	// 计算背包
	packageWeight := (sum + target) / 2
	// 创建动态规划一维数组
	dp := make([]int, packageWeight+1)
	// 动态规划遍历: 先遍历商品再遍历背包
	for _, num := range nums {
		for j := packageWeight; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[packageWeight]
}
