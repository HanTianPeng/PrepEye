package 二分相关

/*
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果 target 存在返回下标，否则返回 -1。
你必须编写一个具有 O(log n) 时间复杂度的算法。
示例 1:
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例 2:
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
提示：
你可以假设 nums 中的所有元素是不重复的。
n 将在 [1, 10000]之间。
nums 的每个元素都将在 [-9999, 9999]之间。
解题思路:
二分查找有三种遍历方式:
开区间: left=-1,right=len; 循环条件:left+1<right; 指针:left=mid,right=mid
闭区间: left=0,right=len-1; 循环条件:left<=right; 指针:left=mid+1,right=mid-1
左闭右开: left=0,right=len; 循环条件:left<right; 指针:left=mid+1,right=mid
*/
// [left, right]; left=0, right=l-1, right=mid-1, left=mid+1
func search(nums []int, target int) int {
	// 统计数组的长度
	l := len(nums)
	// 边界处理
	if l == 0 {
		return -1
	}
	// 二分查找遍历(闭区间)
	left, right := 0, l-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

// [left, right} left=0, right=l, left=mid+1, right=mid
func search1(nums []int, target int) int {
	// 统计数组的长度
	l := len(nums)
	// 边界处理
	if l == 0 {
		return -1
	}
	// 二分查找遍历(左闭右开)
	left, right := 0, l
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid // 左闭右开,所以right=mid是开区间
		} else {
			left = mid + 1
		}
	}
	return -1
}

// (left, right) left=-1, right=l; left=mid, right=mid
func search2(nums []int, target int) int {
	// 统计数组的长度
	l := len(nums)
	// 边界处理
	if l == 0 {
		return -1
	}
	// 二分查找(开区间)
	left, right := -1, l
	for left+1 < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			right = mid
		} else {
			left = mid
		}
	}
	return -1
}
