package main

/*
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回 滑动窗口中的最大值 。
示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
示例 2：
输入：nums = [1], k = 1
输出：[1]
提示：
1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length
解题思路:
1. 保持窗口最多有k个元素,出界就删除队头元素
2. 如果新增的数字大于等于前面的数,那么前面的数肯定不是答案,所以需要从队尾删除
3. 如果新增的数字小于前面的树，那么新增的数字就需要保留
*/
// 239_滑动窗口的最大值: 单调递减队列
func maxSlidingWindow(nums []int, k int) []int {
	// 创建一个结果变量
	ans := make([]int, len(nums)-k+1)
	// 创建一个双端队列变量: 队头队尾都需要删除元素
	queue := make([]int, 0)
	// 遍历
	for i, num := range nums {
		// 维护单调递减性: [5, 4, 1]
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= num {
			queue = queue[:len(queue)-1]
		}
		// 入队列: 队列存数组的下标
		queue = append(queue, i)
		// 从队头出队列: 超出了k的窗口大小
		left := i - k + 1
		if queue[0] < left {
			queue = queue[1:]
		}
		// 计算结果
		if left >= 0 {
			ans[left] = nums[queue[0]]
		}
	}
	return ans
}
