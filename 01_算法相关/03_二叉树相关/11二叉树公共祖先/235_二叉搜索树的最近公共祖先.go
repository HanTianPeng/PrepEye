package main

func lowestCommonAncestorV1(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestorV1(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestorV1(root.Right, p, q)
	} else {
		return root
	}
}

func lowestCommonAncestorV2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
		} else if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}
