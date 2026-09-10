package main

// 寻找重复数: floyd算法 + 初始化head=0
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[nums[0]]
	// 判圈
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// 入口点
	head := 0
	for head != slow {
		head = nums[head]
		slow = nums[slow]
	}
	return slow
}
