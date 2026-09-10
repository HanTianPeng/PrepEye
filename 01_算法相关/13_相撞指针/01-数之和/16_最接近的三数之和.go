package main

import (
	"math"
	"slices"
)

// 最接近的三数之和: 数组排序 + 去重 + 双向指针 + 边界极端判断
func threeSumClosest(nums []int, target int) int {
	// 边界判断: 三数之和,前提条件是数组长度必须大于等于三个;
	if len(nums) < 3 {
		return 0
	}
	// 对数组进行排序
	slices.Sort(nums)

	// 定义一个最近的变量
	minNum := math.MaxInt

	// 定义一个三数之和的结果变量
	ans := 0

	// 两级for遍历
	for first := 0; first < len(nums)-2; first++ {
		// 去重
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 左边界判断
		if nums[first]+nums[first+1]+nums[first+2] >= target {
			diff := nums[first] + nums[first+1] + nums[first+2] - target
			if diff < minNum {
				minNum = diff
				ans = nums[first] + nums[first+1] + nums[first+2]
			}
			continue
		}
		// 右边界判断
		if nums[first]+nums[len(nums)-1]+nums[len(nums)-2] <= target {
			diff := target - nums[first] - nums[len(nums)-1] - nums[len(nums)-2]
			if diff < minNum {
				minNum = diff
				ans = nums[first] + nums[len(nums)-1] + nums[len(nums)-2]
			}
			continue
		}
		// 创建左右指针
		left, right := first+1, len(nums)-1
		for left < right {
			// 计算三个指针的和
			sum := nums[first] + nums[left] + nums[right]
			ans = min(ans, sum)
			if sum < target {
				minNum = min(minNum, target-sum)
				left++
			} else if sum > target {
				minNum = min(minNum, sum-target)
				right--
			} else {
				minNum = 0
				return ans
			}

		}
	}
	return ans
}
