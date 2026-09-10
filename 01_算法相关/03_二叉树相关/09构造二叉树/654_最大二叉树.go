package main

// 最大二叉树
func constructMaximumBinaryTree(nums []int) *TreeNode {
	// 边界处理
	if len(nums) == 0 {
		return nil
	}

	// 求数组的最大值索引
	maxIdx := maxIndex(nums)

	// 边界处理
	if maxIdx == -1 {
		return nil
	}

	// 获取根节点
	root := &TreeNode{Val: nums[maxIdx]}

	// 将数组进行左右切割
	numsLeft := nums[:maxIdx]
	numsRight := nums[maxIdx+1:]

	// 开始构建二叉树
	root.Left = constructMaximumBinaryTree(numsLeft)
	root.Right = constructMaximumBinaryTree(numsRight)
	return root
}

func maxIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	maxIdx, maxNum := 0, nums[0]
	for i, num := range nums {
		if num > maxNum {
			maxNum = num
			maxIdx = i
		}
	}
	return maxIdx
}
