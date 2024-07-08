package main

func postorderTraversal(root *TreeNode) []int {
	var numbers []int
	postorderTraversalArr(root, &numbers)
	return numbers
}

func postorderTraversalArr(root *TreeNode, a *[]int) {
	if root == nil {
		return
	}

	// Recur on left subtree
	postorderTraversalArr(root.Left, a)

	// Recur on right subtree
	postorderTraversalArr(root.Right, a)

	// Deal with the node
	*a = append(*a, root.Val)
}
