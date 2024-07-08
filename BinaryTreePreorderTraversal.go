package main

func preorderTraversal(root *TreeNode) []int {
	var numbers []int
	preorderTraversalArr(root, &numbers)
	return numbers
}

func preorderTraversalArr(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}

	// Deal with the node
	*a = append(*a, root.Val)

	// Recur on left subtree
	preorderTraversalArr(root.Left, a)

	// Recur on right subtree
	preorderTraversalArr(root.Right, a)
}
