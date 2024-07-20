package main

import (
	"fmt"
)

func main() {
	fmt.Println(isHappy(2))
	//00000010100101000001111010011100
	/*	num, err := BinaryToUint32("10100101000001111010011100")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(num)*/
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
