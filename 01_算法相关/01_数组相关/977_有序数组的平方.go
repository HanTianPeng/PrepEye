package main

/*
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
示例 1：
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
示例 2：
输入：nums = [-7,-3,2,3,11]
输出：[4,9,9,49,121]
提示：
1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 已按 非递减顺序 排序
进阶：
请你设计时间复杂度为 O(n) 的算法解决本问题
解题思路:
数组为有序数组，平方过后最大值一定是存在最左边或则最右边;
另外：最大的值在新的数组一定是存在数组的最右边,也就是数组的长度-1
所以可以用双指针
*/
func sortedSquares(nums []int) []int {
	// 创建左右双指针
	left, right, target := 0, len(nums)-1, len(nums)-1
	ans := make([]int, len(nums))
	for left <= right {
		leftNum := nums[left] * nums[left]
		rightNum := nums[right] * nums[right]
		if leftNum > rightNum {
			ans[target] = leftNum
			left++
		} else {
			ans[target] = rightNum
			right--
		}
		target--
	}
	return ans
}
