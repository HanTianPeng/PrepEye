package main

// 摆动序列: 波峰波谷 + 贪心算法 + nums[i+1]-nums[i]
func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	preDiff := 0
	curDiff := 0
	ans := 1
	for i := 0; i < len(nums)-1; i++ {
		curDiff = nums[i+1] - nums[i]
		if (preDiff <= 0 && curDiff > 0) || (preDiff >= 0 && curDiff < 0) {
			preDiff = curDiff
			ans++
		}
	}
	return ans
}
