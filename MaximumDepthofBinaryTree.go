package main

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var lMax = maxDepth(root.Left)
	var rMax = maxDepth(root.Right)

	if lMax > rMax {
		return lMax + 1
	} else {
		return rMax + 1
	}
}
