package main

/*
nums: [1, 2, 3, 1]
输出: 4
解题思路:
如果打劫第i家，那么能打劫的最大金额就是:
一定要打劫第i家
f[i] = f[i-2] + nums[i]
一定不要打劫第i家
f[i] = f[i-1]
两者的最大值就是打劫第i家的最大金额
由于防止边界越界,所以创建的数组的长度默认+2
*/

func rob(nums []int) int {
	l := len(nums)
	f := make([]int, l+2)
	// 动态规划遍历
	for i, num := range nums {
		f[i+2] = max(f[i+1], f[i]+num)
	}
	return f[l+1]
}

func robV2(nums []int) int {
	f0, f1 := 0, 0
	for _, num := range nums {
		f0, f1 = f1, max(f0+num, f1)
	}
	return f1
}
