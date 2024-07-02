package main

import (
	"fmt"
)

func main() {
	var array = []int{886, 729}
	fmt.Println(maxProfit(array))
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
