package main

import (
	"fmt"
)

func main() {
	input := "A man, a plan, a canal: Panama"

	a := isPalindrome(input)
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
