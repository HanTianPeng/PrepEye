package main

import "fmt"

// 当前节点的值-1如果存在,则说明更长的一定是在后面
func longestConsecutive(nums []int) int {
	// 创建一个map,保证检索的时候为o(1)的时间复杂度
	mp := make(map[int]bool)
	for _, n := range nums {
		mp[n] = true
	}
	// 遍历起点
	ans := 0
	for n := range mp {
		// 如果当前节点不是最左边的节点,则相当于更长;直接抛弃当前节点
		if mp[n-1] {
			continue
		}
		m := n + 1
		for mp[m] {
			m++
		}
		if m-n > ans {
			ans = m - n
		}
	}
	return ans
}

func main() {
	res := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	fmt.Println(res)
}
