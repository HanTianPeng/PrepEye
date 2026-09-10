package main

/*
杨辉三角
1 -> 1
2 -> [[1], [1, 1]]
3 -> [[1], [1, 1], [1, 2,1]]
4 -> [[1], [1, 1], [1, 2, 1], [1, 3, 3, 1]]
f[i][0] = f[i][i] = 1
f[i][j] = f[i-1][j-1] + f[i-1][j]
*/
func generate(numRows int) [][]int {
	// 定义一个返回变量
	ans := make([][]int, numRows)
	// 边界处理
	if numRows <= 0 {
		return ans
	}
	// 将numRows=1数据先初始化
	ans[0] = []int{1}
	// 动态规划计算numRows=2,3等之后的数据
	for i := 1; i < numRows; i++ {
		// 定义一个上一层的数组变量
		preNums := ans[i-1]
		// 定义当前层的数组
		var res []int
		// 初始化最初数字
		preNum := 0
		for _, curNum := range preNums {
			res = append(res, curNum+preNum)
			preNum = curNum
		}
		// 最后一个元素也需要加进去
		res = append(res, preNums[len(preNums)-1])
		ans[i] = res
	}
	return ans
}

func generateV2(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0], ans[i][i] = 1, 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}
