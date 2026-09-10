package main

/*
解题思路:
[1, 2, 3, 4] => 5
[-1, 0, 3, 4, 5] => 1
[2, 2, 1, 1, 3] => 4
*/
// 41_缺失的第一个正数: 换位置,每个数组的值都到自己该去的位置,只有一个位置不符合
func firstMissingPositive(nums []int) int {
	l := len(nums)

	// 遍历
	for i := range l {
		// nums[i]必须是在[1, l]之间;因为最后有l个同学,理论上l个同学的学号应该是1,2,3...l;否则认为无效学号
		// l个同学的学号可能出现 1, 1, 2, 3重复元素;只要出现重复，也不进行换位置
		for nums[i] >= 1 && nums[i] <= l && nums[nums[i]-1] != nums[i] {
			// j: nums[i] = 4,这就代表4号学生应该回归到索引3的位置;所以j是nums[i]的学号对应的索引
			j := nums[i] - 1
			nums[j], nums[i] = nums[i], nums[j]
		}
	}

	// 理论上l个同学的学号是1,2,3...l;如果位置不对,证明缺少了就是谁的学号
	for i := range l {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return l + 1
}
