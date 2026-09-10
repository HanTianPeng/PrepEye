package main

// 将有序数组转换为二叉搜索树: 根->左->右
func sortedArrayToBST(nums []int) *TreeNode {
	// 递归边界处理
	if len(nums) == 0 {
		return nil
	}
	// 将中点作为根节点
	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	// 构建左子树和右子树
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
