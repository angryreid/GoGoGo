package main

import (
	"fmt"
	queue2 "go-learing/mid/queue"
)

func main() {
	var queue = queue2.Queue{}
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.IsEmpty())
	queue.Pop()
	fmt.Println(queue.IsEmpty())
	queue.Pop()
	fmt.Println(queue.IsEmpty())
}
