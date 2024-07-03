package main

import (
	"fmt"
)

func main() {
	var nums = []int{1}
	a := singleNumber(nums)
	fmt.Println(a)
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
