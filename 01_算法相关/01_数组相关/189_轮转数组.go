package main

/*
解题思路:
1 -> 2 -> 3 -> 4 -> 5 -> 6   k = 3
4 -> 5 -> 6 -> 1 -> 2 -> 3
首先k会发生越界数组的长度,所以求余
交换位置思想: 总共发生换位置的次数为数组的长度;
从1开始:
1 -> 4
4 -> 1
这一轮结束
2 -> 5
5 -> 2
这一轮再次结束
3 -> 6
6 -> 3
最后一轮结束
*/
func rotate(nums []int, k int) {
	// 计算数组的长度
	l := len(nums)
	// 边界判断
	if l == 0 || l == 1 {
		return
	}
	// 如果数组的长度是3.轮转3个位置;就是还原了
	k = k % l
	// 空间复杂度为O(1),那么就是原地轮换，开始从0号索引位置的元素进行跳位置;总共跳转l次
	startIdx, count := 0, 0
	for count < l {
		// 0号索引移动到3号索引, 6号索引移动到2号索引(6+3)%7 = 2
		curIdx, curVal := startIdx, nums[startIdx]
		for {
			nxtIdx := (curIdx + k) % l
			nxtVal := nums[nxtIdx]
			nums[nxtIdx] = curVal
			curIdx = nxtIdx
			curVal = nxtVal
			count++
			if curIdx == startIdx {
				// 一轮换位置周期闭环
				break
			}
		}
		startIdx++
	}
}

/*
解题思路:
1 -> 2 -> 3 -> 4 -> 5 -> 6    k = 3
第一步: 反转整个数组
6 -> 5 -> 4 -> 3 -> 2 -> 1
第二步: 按照k进行左右两半反转
4 -> 5 -> 6 -> 3 -> 2 -> 1  先反转左半段
4 -> 5 -> 6 -> 1 -> 2 -> 3  后反转右半段
*/
func rotate2(nums []int, k int) {
	// k的求余处理
	k = k % len(nums)
	// 首先反转整个数组
	reverseList(nums)
	reverseList(nums[:k])
	reverseList(nums[k:])
}

/*
数组的反转:
解题思路: 镜像数组, 左右指针互换位置；需要互换位置的次数是长度的一半次
*/
func reverseList(nums []int) {
	// 反转数组其实就是镜像数组;所以需要进行镜像换位置l/2次
	count, leftIdx, rightIdx := len(nums)/2, 0, len(nums)-1
	for count > 0 {
		temp := nums[rightIdx]
		nums[rightIdx] = nums[leftIdx]
		nums[leftIdx] = temp
		count--
		leftIdx++
		rightIdx--
	}
}
