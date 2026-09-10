package main

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
解题思路:
假设最后一步为爬1个台阶: f[i]就有f[i - 1] 个方法
假设最后一步为爬2个台阶: f[i]就有f[i - 2]个方法
所以得出结论: f[i] = f[i - 1] + f[i - 2]
*/
func climbStairs(n int) int {
	f0, f1 := 1, 1
	for i := 2; i <= n; i++ {
		temp := f1
		f1 += f0
		f0 = temp
	}
	return f1
}

func climbStairs2(n int) int {
	// 1. 确定dp数组含义、下标含义: dp[i] 到达第n个台阶有多少种办法
	// dp[i] = dp[i - 1] + dp[i - 2]

	// 2. 递推公式:最后是一步; 所以 dp[i] = dp[i - 1]; 最后是两步: 所以 dp[i] = dp[i - 2]; 所以dp[i] 有dp[i - 1] + dp[i - 2]中方法

	// 3. 初始化: dp[1] = 1 dp[2] = 2 ; 为什么不初始化dp[0]=0呢;因为要求n大于等于1; 所以需要边界判断
	if n <= 2 {
		return n
	}

	dp := make([]int, n+1)
	dp[1], dp[2] = 1, 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
