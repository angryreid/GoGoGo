package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func CreateTree(value int) *Node {
	return &Node{Value: value}
}

func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func (node *Node) TraverseFunc(f func(*Node)) {
	if node == nil {
		return
	}

	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *Node) TraverseWithChannel() chan *Node {
	sender := make(chan *Node)
	go func() {
		node.TraverseFunc(func(node *Node) {
			sender <- node
		})
		close(sender)
	}()

	return sender
}
