package main

/*
给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
示例 1：
1 2 3
8 9 4
7 6 5
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
解题思路:
这个是矩阵的遍历: 和二分查找一样，for循环遍历的边界固定，左闭右开
第一轮是遍历: 1, 2
第二轮是遍历: 3, 4
第三轮是遍历: 5, 6
第四轮是遍历: 7, 8
这是一圈遍历;
遍历的圈数等于 n / 2
如果n不能被整除,那么中心位置一定是 ans[n/2][n/2] n*n
*/
func generateMatrix(n int) [][]int {
	// 初始化一个二元数组
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}

	// 循环的次数
	loop := n / 2
	// 中心点
	center := n / 2
	// 计数器
	count := 1
	// 偏移量
	offset := 1
	// 遍历的第一圈
	starX, starY := 0, 0
	for loop > 0 {
		x, y := starX, starY
		// 最顶部的一行，行数不变，列数+1; 从左至右(左闭右开)  1, 2,
		for y = starY; y < n-offset; y++ {
			ans[starX][y] = count
			count++
		}

		// 最右侧的一列, 列数不变，行数+1; 从上到下(左闭右开)  3, 4
		for x = starX; x < n-offset; x++ {
			ans[x][y] = count
			count++
		}

		// 最底部的一行, 行数不变，列数-1; 从右到左(左闭右开)  5, 6
		for ; y > starY; y-- {
			ans[x][y] = count
			count++
		}

		// 最左边的一行,列数不变，行数-1; 从下到上(左闭右开)  7, 8
		for ; x > starX; x-- {
			ans[x][y] = count
			count++
		}
		offset++
		loop--
		starX++
		starY++
	}
	if n%2 == 1 {
		ans[center][center] = n * n
	}
	return ans
}
