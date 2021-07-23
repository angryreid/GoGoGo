package main

import "fmt"

func main() {
	msg := make(chan string)

	go func() {
		goRunStr := "Go Routine running."
		fmt.Println(goRunStr)
		msg <- goRunStr
	}()
	fmt.Println("Hello go one more time!")
	fmt.Println(<-msg)
}
