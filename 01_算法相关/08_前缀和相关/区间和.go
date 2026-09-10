package main

/*
解题思路:
前缀和在涉及计算区间和的问题时fei
*/
// 区间和: 前缀和 + n+1数组 + (arr[0-right]-arr[0-left])
func ArrageSum(nums []int, left, right int) int {
	// 定义一个变量存储前缀和
	ans := make([]int, len(nums)+1)
	for i, num := range nums {
		if i == 0 {
			ans[i] = num
		} else {
			ans[i] = num + ans[i-1]
		}
	}
	// [left, right] = [0, right] - [0, left-1]
	return ans[right] - ans[left-1]
}
