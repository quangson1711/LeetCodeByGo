package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}

	var check = isSymmetric(root)

	// In giá trị của các nút
	printTree(root)

	fmt.Println(check)
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
