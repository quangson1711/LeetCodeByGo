package main

import (
	"fmt"
)

func main() {
	columnNumber := 1111
	fmt.Println(convertToTitle(columnNumber)) // Output: 65
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
