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
重要部分:
如果使用一维数组: 物品遍历for循环放在外层,背包遍历for循环放在内层,且内层循环倒序遍历;
-> 倒序遍历是为了保证物品i只被放入一次;正序dp[1], dp[2]不就相当于正序遍历就被重复计算了吗
*/
func maxValue(goodsWeightList, goodsValueList []int, packageWeight int) int {
	// 初始化一个二维数组
	goodsLen := len(goodsWeightList)
	dp := make([][]int, goodsLen)
	for i := range dp {
		dp[i] = make([]int, packageWeight+1)
		dp[i][0] = 0
	}

	// 初始化边界数据
	for j := 1; j >= goodsWeightList[0] && j <= packageWeight; j++ {
		dp[0][j] = goodsValueList[0]
	}
	for j := 1; j < goodsWeightList[0] && j <= packageWeight; j++ {
		dp[0][j] = 0
	}

	// 动态规划-先遍历商品
	for i := 1; i < goodsLen; i++ {
		// 再遍历背包
		for j := 1; j <= packageWeight; j++ {
			// 如果当前背包的重量小于商品i的重量
			if j < goodsWeightList[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-goodsWeightList[i]]+goodsValueList[i])
			}
		}
	}
	return dp[goodsLen-1][packageWeight]
}

func maxValueV2(goodsWeightList, goodsValueList []int, packageWeight int) int {
	// 商品的总数量
	goodsLen := len(goodsWeightList)
	// 初始化背包
	dp := make([]int, packageWeight+1)
	// 初始化数据

	// 动态规划-先遍历商品
	for i := 0; i < goodsLen; i++ {
		for j := packageWeight; j >= goodsWeightList[i]; j-- {
			dp[j] = max(dp[j], dp[j-goodsWeightList[i]]+goodsValueList[i])
		}
	}
	return dp[packageWeight]
}

func main() {
	goodsWeightList := []int{1, 3, 4}
	goodsValueList := []int{15, 20, 30}
	packageWeight := 4
	res := maxValue(goodsWeightList, goodsValueList, packageWeight)
	fmt.Println("res:", res)
	resv2 := maxValueV2(goodsWeightList, goodsValueList, packageWeight)
	fmt.Println("resv2:", resv2)
}
