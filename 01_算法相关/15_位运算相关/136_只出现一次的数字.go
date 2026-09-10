package main

/*
解题思路:
位运算:
 1. 任何数和自己异或结果都为0
 2. 任何数和0异或都不变
*/
func singleNumber(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
