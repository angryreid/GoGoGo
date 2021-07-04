package main

import (
	"fmt"
	"go-learing/mid/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 1}
	fmt.Println(root) // {1 nil nil}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 2}
	root.Left.Right = new(tree.Node)
	fmt.Println(root)
	root.Left.Right.Print() // 0
	root.Left.Right.SetValue(100)
	root.Left.Right.Print() // 100
	fmt.Println("Tree Traverse:")
	root.Traverse()
}
