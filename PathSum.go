package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	targetSum = targetSum - root.Val

	if root.Left == nil {
		return hasPathSum(root.Right, targetSum)
	}

	if root.Right == nil {
		return hasPathSum(root.Left, targetSum)
	}

	lCheck := hasPathSum(root.Left, targetSum)
	rCheck := hasPathSum(root.Right, targetSum)
	if lCheck || rCheck {
		return true
	}
	return false
}
