package main

import "slices"

// 数组排序 + 双向指针 + 去重 + 边界极端判断 + 双指针过程去重
func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	if len(nums) < 3 {
		return ans
	}
	// 数组排序
	slices.Sort(nums)
	for first := 0; first < len(nums)-2; first++ {
		// 跳过重复的元素
		if first > 0 && nums[first-1] == nums[first] {
			continue
		}
		// 跳过和大于0的直接break
		if nums[first]+nums[first+1]+nums[first+2] > 0 {
			break
		}
		// 跳过和小于0的直接break
		if nums[first]+nums[len(nums)-1]+nums[len(nums)-2] < 0 {
			continue
		}
		// 创建双指针
		left, right := first+1, len(nums)-1
		for left < right {
			sum := nums[first] + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{nums[first], nums[left], nums[right]})
				left++
				right--
				// 去除过程中的重复元素
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// 去除过程中的重复元素
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}
