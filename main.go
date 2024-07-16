package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 2, 1, 1, 1, 2}
	fmt.Println(majorityElement(nums))
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
