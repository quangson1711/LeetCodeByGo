package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	printTree(root)

	numbers := postorderTraversal(root)
	fmt.Println(numbers)
}

// Hàm in giá trị của các nút
func printTree(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("Node value: %d\n", node.Val)
	printTree(node.Left)
	printTree(node.Right)
}
