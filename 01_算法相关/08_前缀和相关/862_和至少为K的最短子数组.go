package main

// 862_和至少为K的最短子数组: 前缀和 + 单调递增双端队列
func shortestSubArray(nums []int, k int) int {
	l := len(nums)
	// 创建一个前缀和数组
	s := make([]int, l+1)
	for i, num := range nums {
		s[i+1] = s[i] + num
	}
	// 创建一个单调递增的双端队列
	ans := l + 1
	queue := make([]int, 0)
	for i, js := range s {
		// 如果js-s[i1]>=k,那么相当于符合目标的结果,但是需要继续向ss[i2]寻找最优解
		for len(queue) > 0 && js-s[queue[0]] >= k {
			ans = min(ans, i-queue[0])
			queue = queue[1:]
		}
		// 如果i1<i2,s[i1]>=s[i2];那么j-s[i1]<=j-s[i2];如果j-s[i1]都符合k,那么j-s[i2]更符合k,所以直接丢弃i1
		for len(queue) > 0 && s[queue[len(queue)-1]] >= js {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
	}
	if ans > l {
		return -1
	}
	return ans
}
