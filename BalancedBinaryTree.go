package main

func isBalanced(root *TreeNode) bool {
	if check := isBalancedInt(root); check >= 0 {
		return true
	}
	return false
}

func isBalancedInt(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := isBalancedInt(root.Left)
	if lh == -1 {
		return -1
	}

	rh := isBalancedInt(root.Right)
	if rh == -1 {
		return -1
	}

	if abs := absInt(lh - rh); abs > 1 {
		return -1
	} else {
		return 1 + maxInt(lh, rh)
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
