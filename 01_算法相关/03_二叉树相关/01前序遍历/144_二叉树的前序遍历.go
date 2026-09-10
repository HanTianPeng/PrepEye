package main

/*
树:
  1. 一种非线性的数据结构,分层存储
  2. 仅有唯一一个根节点,没有节点则为空树
  3. 除根节点外,每个节点都有并仅有唯一一个父节点
  4. 节点间不能形成闭环
树的常见概念:
  1. 拥有相同父节点的节点,相互称为兄弟节点
  2. 节点的深度: 从根节点到该节点所经历的边的个数
  3. 节点的高度: 节点到叶节点的最长路径
  4. 树的高度: 根节点的高度
二叉树: 最多仅有两个子节点的树(最多能分成两个叉的树)
平衡二叉树: 二叉树中,每个节点的左右子树的高度相差不能大于1
满二叉树: 除了叶节点外每一个节点都有左右子叶且叶子节点都处在最底层的二叉树
完全二叉树: 深度为h,除第h层外,其他各层(1--h-1)的节点数都达到最大个数,第h层所有的节点都连续集中在最左边
       A
   B       C
D    E   F    G
前序遍历: 根->左->右  A->B->D->E->C->F->G
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
二叉树的前序遍历: 根 -> 左 -> 右
*/
func preorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	ans = append(ans, root.Val)
	ans = append(ans, preorderTraversal(root.Left)...)
	ans = append(ans, preorderTraversal(root.Right)...)
	return ans
}

func preorderTraversal2(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 {
		// 从栈的顶部弹出一个节点
		node := stack[len(stack)-1]
		// 将栈的大小进行缩容
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ans
}
