package main

import "fmt"

/*
对于多重背包，我在力扣上还没发现对应的题目，所以这里就做一下简单介绍，大家大概了解一下。
有N种物品和一个容量为V 的背包。第i种物品最多有Mi件可用，每件耗费的空间是Ci ，价值是Wi 。求解将哪些物品装入背包可使这些物品的耗费的空间 总和不超过背包容量，且价值总和最大。
多重背包和01背包是非常像的， 为什么和01背包像呢？
每件物品最多有Mi件可用，把Mi件摊开，其实就是一个01背包问题了。
例如：
背包最大重量为10。
物品为：
重量	价值	数量
物品0	1	15	2
物品1	3	20	3
物品2	4	30	2
问背包能背的物品最大价值是多少？
和如下情况有区别么？
重量	价值	数量
物品0	1	15	1
物品0	1	15	1
物品1	3	20	1
物品1	3	20	1
物品1	3	20	1
物品2	4	30	1
物品2	4	30	1
毫无区别，这就转成了一个01背包问题了，且每个物品只用一次。
*/
// 动态规划二维数组: 先遍历背包再遍历商品
func maxValueMultiNum(goodsWeightList, goodsValueList, goodsNumList []int, packageWeight int) int {
	// 计算商品的总数
	goodsNum := 0
	for _, num := range goodsNumList {
		goodsNum += num
	}
	// 创建动态规划二维数组
	dp := make([][]int, goodsNum+1)
	for i := range dp {
		dp[i] = make([]int, packageWeight+1)
		dp[i][0] = 0
	}
	// 初始化
	for j := 1; j <= packageWeight; j++ {
		dp[0][j] = 0
	}
	// 动态规划遍历: 先遍历背包
	for j := 1; j <= packageWeight; j++ {
		total := 1
		for i := 0; i < len(goodsWeightList); i++ {
			for k := range goodsNumList[i] {
				if j >= goodsWeightList[i]*(k+1) {
					dp[total][j] = max(dp[total-1][j], dp[total-1][j-goodsWeightList[i]]+goodsValueList[i])
				} else {
					dp[total][j] = dp[total-1][j]
				}
				total += 1
			}
		}
	}
	return dp[goodsNum][packageWeight]
}

// 一维数组: 先遍历商品再遍历背包
func maxValueMultiNumV2(goodsWeightList, goodsValueList, goodsNumList []int, packageWeight int) int {
	// 创建化动态规划一维数组
	dp := make([]int, packageWeight+1)
	// 初始化
	dp[0] = 0
	// 动态规划遍历: 先遍历商品再遍历背包
	for i := 0; i < len(goodsWeightList); i++ {
		for j := packageWeight; j >= goodsWeightList[i]; j-- {
			for k := range goodsNumList[i] {
				if j >= goodsWeightList[i]*(k+1) {
					dp[j] = max(dp[j], dp[j-goodsWeightList[i]*(k+1)]+goodsValueList[i]*(k+1))
				}
			}
		}
	}
	return dp[packageWeight]
}

// 一维数组: 先遍历商品再遍历背包
func maxValueMultiNumV3(goodsWeightList, goodsValueList, goodsNumList []int, packageWeight int) int {
	// 创建化动态规划一维数组
	dp := make([]int, packageWeight+1)
	// 初始化
	dp[0] = 0
	// 动态规划遍历: 先遍历商品再遍历背包
	for i := 0; i < len(goodsWeightList); i++ {
		for k := range goodsNumList[i] {
			for j := packageWeight; j >= goodsWeightList[i]*(k+1); j-- {
				// dp[1] = 15 dp[3]
				temp := max(dp[j], dp[j-goodsWeightList[i]*(k+1)]+goodsValueList[i])
				fmt.Println(i, k+1, j, dp[j], j-goodsWeightList[i]*(k+1), dp[j-goodsWeightList[i]*(k+1)], goodsValueList[i], temp)
				dp[j] = temp
			}
		}
	}
	//fmt.Println(dp)
	return dp[packageWeight]
}

func main() {
	goodsWeightList := []int{1, 3, 4}
	goodsValueList := []int{15, 20, 30}
	goodsNumList := []int{2, 3, 2}
	//res1 := maxValueMultiNumV3(goodsWeightList, goodsValueList, goodsNumList, 1)
	//fmt.Println("res1:", res1) // 15
	res2 := maxValueMultiNumV3(goodsWeightList, goodsValueList, goodsNumList, 2)
	fmt.Println("res2:", res2) // 30
	//res3 := maxValueMultiNumV3(goodsWeightList, goodsValueList, goodsNumList, 3)
	//fmt.Println("res3:", res3) // 30
	//res4 := maxValueMultiNumV3(goodsWeightList, goodsValueList, goodsNumList, 4)
	//fmt.Println("res4:", res4) // 35
	//res5 := maxValueMultiNumV3(goodsWeightList, goodsValueList, goodsNumList, 5)
	//fmt.Println("res5:", res5) // 50
	//res6 := maxValueMultiNumV3(goodsWeightList, goodsValueList, goodsNumList, 6)
	//fmt.Println("res6:", res6) // 60 4 + 1 + 1
	//res7 := maxValueMultiNumV3(goodsWeightList, goodsValueList, goodsNumList, 7)
	//fmt.Println("res7:", res7) // 60 3 + 3 + 1 = 20 + 20 + 15 = 55  4 + 1 + 1 => 60
	//res8 := maxValueMultiNumV3(goodsWeightList, goodsValueList, goodsNumList, 8)
	//fmt.Println("res8:", res8) // 3 + 3 + 1 + 1 => 70
	//res9 := maxValueMultiNumV3(goodsWeightList, goodsValueList, goodsNumList, 9)
	//fmt.Println("res9:", res9) // 4 + 3 + 1 + 1 => 80
	//res10 := maxValueMultiNumV3(goodsWeightList, goodsValueList, goodsNumList, 10)
	//fmt.Println("res10:", res10) // 4 + 4 + 1 + 1 => 90
}
