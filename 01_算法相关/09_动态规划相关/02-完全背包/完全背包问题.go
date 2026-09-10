package main

import "fmt"

/*
解题思路:
goodsWeightList: [1, 3, 4]
goodsValueList: [15, 20, 30]
packageWeight: 4
状态转换方程式:
dp[i][j]: 表示i个商品，重量为j的背包能装的最大价值
1.选商品i,那产生的价值就是:dp[i-1][j-goodsWeightList[i]] + goodsValueList[i]
2.不选商品i,那产生的价值就是:dp[i-1][j]
既然最终是为了得到i个商品钟,用重量为j的背包进行装商品,不重复装,能装的最大价值就是: 0/1背包
dp[i][j] = max(dp[i-1][j], dp[i-1][j-goodsWeightList[i]] + goodsValueList[i])
dp[i][0] = 0
dp[0][j]:
1. 当背包j的重量大于等于商品0的重量，则价值为商品0的价值
2. 当背包j的重量小于商品0的重量,则价值为0

完全背包: 与[0/1背包的区别在于商品i是否可以重复选择]
dp[i][j] = max(dp[i-1][j], dp[i][j-goodsWeightList[i] + goodsValueList[i]])

如果求组合数: 不讲究遍历顺序: 所以一定要进行->先遍历商品,再遍历背包
如果求排列数; 需讲究遍历顺序: 所以一定要进行->先遍历背包,再遍历商品
*/
// 完全背包
func maxValueMulti(goodsWeightList, goodsValueList []int, packageWeight int) int {
	// 1. dp[i][j]的定义: i个商品，使用重量为j的背包,能装起的最大价值;单个商品可以重复
	goodsLen := len(goodsWeightList)
	dp := make([][]int, goodsLen)
	for i := range dp {
		dp[i] = make([]int, packageWeight+1)
	}
	// 2. 动态规划方程式: dp[i][j] = max(dp[i-1][j], dp[i][j-goodsWeightList[i]] + goodsValueList[i])
	// 3. 初始化:当背包j为0的时候,dp[i][0] = 0; 当商品i为0的时候,dp[0][j]就分成两种情况
	for j := goodsWeightList[0]; j <= packageWeight; j++ {
		// 针对于背包>=商品0的重量时候,能多几次商品0就装几次
		// 针对于背包<商品0的重量时候,dp[0][j]=0;go语言不需要额外初始化
		dp[0][j] = dp[0][j-goodsWeightList[0]] + goodsValueList[0]
	}
	// 动态规划遍历:先遍历商品
	for i := 1; i < goodsLen; i++ {
		// 再遍历背包
		for j := 1; j <= packageWeight; j++ {
			// 如果商品的重量大于背包的重量
			if j < goodsWeightList[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-goodsWeightList[i]]+goodsValueList[i])
			}
		}
	}
	return dp[goodsLen-1][packageWeight]
}

func main() {
	goodsWeightList := []int{1, 3, 4}
	goodsValueList := []int{15, 20, 30}
	packageWeight := 4
	res := maxValueMulti(goodsWeightList, goodsValueList, packageWeight)
	fmt.Println("res:", res)
	resv2 := maxValueMulti(goodsWeightList, goodsValueList, packageWeight)
	fmt.Println("resv2:", resv2)
}
