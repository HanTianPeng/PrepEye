package main

/*
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。
示例 2：
输入：target = 4, nums = [1,4,4]
输出：1
示例 3：
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0
提示：
1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 104
解题思路:
滑动窗口: 双指针
窗口如何定义: sum大于等于target
左指针如何移动: 当窗口大小第一次达到大于等于target,left++,窗口缩小
右指针如何移动: 当窗口大小小于target,right++,窗口放大
*/
// 209_长度最小的子数组: 滑动窗口 + 前提条件数组为正整数
func minSubArrayLen(target int, nums []int) int {
	// 创建结果变量: +1是因为存在不符合预期的时候方便进行边界判断
	ans := len(nums) + 1
	// 创建滑动窗口变量
	left, sum := 0, 0
	for right, num := range nums {
		sum += num
		// 滑动窗口的判断条件: 前提条件数组里面元素都是正整数,不然左指针缩容的时候,遇到一个负数,就不具备单调性,这个时候就需要用前缀和方法来解决
		for sum >= target {
			ans = min(ans, right-left+1)
			sum -= nums[left]
			left += 1
		}
	}
	// 边界判断
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}
