package main

/*
[1, 2, 3, 4]
f[i] = left[i - 1] * right[i + 1]
left[i] = left[i - 1] * nums[i - 1]
right[i] = right[i + 1] * nums[i + 1]
left[0] = 1 => 1乘以任何数等于任何数
*/
func productExceptSelf(nums []int) []int {
	leftNums := make([]int, len(nums))
	leftNums[0] = 1
	for i := 1; i < len(nums); i++ {
		leftNums[i] = leftNums[i-1] * nums[i-1]
	}

	rightNUms := make([]int, len(nums))
	rightNUms[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightNUms[i] = rightNUms[i+1] * nums[i+1]
	}

	ans := make([]int, len(nums))
	for i, leftNum := range leftNums {
		ans[i] = leftNum * rightNUms[i]
	}
	return ans
}

func productExceptSelfV2(nums []int) []int {
	// 优化空间复杂度;
	ans := make([]int, len(nums))
	ans[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		ans[i] = ans[i+1] * nums[i+1]
	}

	// 右边 * 左边
	pre := 1
	for i := 0; i < len(nums); i++ {
		ans[i] = ans[i] * pre
		pre = pre * nums[i]
	}
	return ans
}
