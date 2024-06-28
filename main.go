package main

import (
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	var check = isBalanced(root)

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
