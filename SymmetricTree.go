package main

func isSymmetricTwoTree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a != nil && b != nil {
		if a.Val == b.Val {
			if isSymmetricTwoTree(a.Left, b.Right) {
				if isSymmetricTwoTree(a.Right, b.Left) {
					return true
				}
			}
		}
	}
	return false
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricTwoTree(root.Left, root.Right)
}
