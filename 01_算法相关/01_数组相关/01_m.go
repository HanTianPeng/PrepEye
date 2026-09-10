package main


// 有序递增数组,找到>=num的最小索引
func findMostNumIdx(nums []int, num int) int {
	// 边界处理
	if len(nums) == 0 {
		return -1
	}
	// 二分查找
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= num {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

// 有序递增数组,找到<=num的最大索引
func findLessNumIdx(nums []int, num int) int {
	// 边界处理
	if len(nums) == 0 {
		return -1
	}
	// 二分查找
	l, r := 0, len(nums)-1
	ans := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] <= num {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}

// 局部最小:
func oneMixIdx(nums []int) int {
	// 边界处理
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return -1
	}
	if len(nums) == 2 {
		if nums[0] < nums[1] {
			return 0
		}
		return 1
	}
	if string()
}
