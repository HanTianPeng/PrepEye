package main

import "slices"

// 四数之和: 数组排序 + 去重 + 两层for循环 + 相向指针 + 边界极端判断 + 双指针过程去重
func fourSum(nums []int, target int) [][]int {
	// 定义一个结果变量
	ans := make([][]int, 0)
	// 边界判断
	if len(nums) < 4 {
		return ans
	}
	// 数组排序
	slices.Sort(nums)
	// 遍历
	for first := 0; first < len(nums)-3; first++ {
		// first作为第一个数不能重复
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// 左边界判断
		if nums[first]+nums[first+1]+nums[first+2]+nums[first+3] > target {
			continue
		}
		// 右边界判断
		if nums[first]+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target {
			continue
		}
		for second := first + 1; second < len(nums)-2; second++ {
			// second作为第二个数不能重复
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 左边界判断
			if nums[first]+nums[second]+nums[second+1]+nums[second+2] > target {
				continue
			}
			// 右边界判断
			if nums[first]+nums[second]+nums[len(nums)-1]+nums[len(nums)-2] < target {
				continue
			}
			// 创建相向双指针
			left, right := second+1, len(nums)-1
			for left < right {
				sum := nums[first] + nums[second] + nums[left] + nums[right]
				if sum < target {
					left++
				} else if sum > target {
					right--
				} else {
					ans = append(ans, []int{nums[first], nums[second], nums[left], nums[right]})
					left++
					right--
					// 剔除过程中的重复元素
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					// 剔除过程中的重复元素
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				}
			}
		}
	}
	return ans
}
