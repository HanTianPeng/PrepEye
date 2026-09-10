package main

// 移动零: 单指针记录非0的索引 + 交换位置
func moveZeroes(nums []int) {
	i0 := 0
	for i, x := range nums {
		if x != 0 {
			nums[i], nums[i0] = nums[i0], x
			i0++
		}
	}
}

// 283-移动0: 单指针
func moveZeroes2(nums []int) {
	// 创建一个指针,始终保留最左边的0的索引位置
	var pointer int
	for i := 0; i < len(nums); i++ {
		// 边界处理
		if i == 0 {
			if nums[i] == 0 {
				pointer = i
			}
			continue
		}
		// 遍历的元素是否为0,如果是0,需要判断前面元素是否为0来决定指针的滑动
		if nums[i] == 0 && nums[i-1] != 0 {
			pointer = i
			continue
		}
		// 遍历的元素如果不为0,则需要判断前面一个元素是否为0来决定互换位置
		if nums[i] != 0 && nums[i-1] == 0 {
			nums[pointer] = nums[i]
			nums[i] = 0
			pointer++
		}
	}
}

// 栈 + 单指针
// 单指针记录当前不为0的位置,将不为零的元素移动到单指针的位置
// 剩余元素直接clear;这样剩余元素默认都为0
func moveZeroes3(nums []int) {
	stackNum := 0
	for _, x := range nums {
		if x != 0 {
			nums[stackNum] = x
			stackNum++
		}
	}
	clear(nums[stackNum:])
}
