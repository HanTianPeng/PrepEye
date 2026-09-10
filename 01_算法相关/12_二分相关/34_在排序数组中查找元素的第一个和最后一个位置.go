package 二分相关

/*
给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
示例 1：
输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]
示例 2：
输入：nums = [5,7,7,8,8,10], target = 6
输出：[-1,-1]
示例 3：
输入：nums = [], target = 0
输出：[-1,-1]
提示：
0 <= nums.length <= 105
-109 <= nums[i] <= 109
nums 是一个非递减数组
-109 <= target <= 109
*/
func searchRange(nums []int, target int) []int {
	// 统计数组的长度
	l := len(nums)
	// 边界判断
	if l == 0 {
		return []int{-1, -1}
	}
	// 二分查找遍历:闭区间
	minNum, isFirst, maxNum := -1, true, -1
	left, right, mLeft, mRight := 0, l-1, 0, l-1

	// 寻找最大值
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			if isFirst {
				maxNum = mid
				mLeft = mid + 1
				mRight = right
				isFirst = false
			}
			minNum = mid
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// 寻找最小值
	for mLeft <= mRight {
		mid := mLeft + (mRight-mLeft)/2
		if nums[mid] == target {
			maxNum = mid
			mLeft = mid + 1
		} else if nums[mid] < target {
			mLeft = mid + 1
		} else {
			mRight = mid - 1
		}
	}

	return []int{minNum, maxNum}
}

func searchRange2(nums []int, target int) []int {
	minNum := findMin(nums, target)
	maxNum := findMax(nums, target)
	return []int{minNum, maxNum}
}

func findMin(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			ans = mid
			r = mid - 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func findMax(nums []int, target int) int {
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			ans = mid
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func searchRange3(nums []int, target int) []int {
	// 获取最小索引
	minNum := binarySearch(nums, target)
	// minNum = len(nums)这个代表left指针移动到最右边还是没有找到
	if minNum == len(nums) || nums[minNum] != target {
		return []int{-1, -1}
	}
	// 获取最大索引
	maxNum := binarySearch(nums, target+1) - 1
	return []int{minNum, maxNum}

}

// 查找元素的最左索引，查找不到返回left指针
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
