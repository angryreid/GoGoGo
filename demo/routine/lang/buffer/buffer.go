package main

import (
	"fmt"
)

func sample(message chan string) {
	message <- "Hello 1"
	message <- "Hello 2"
	message <- "Hello 3"
	close(message) // if no this line, will be fatal error, deadlock
	//message <- "Hello 4"
}

func main() {
	var message = make(chan string, 3)
	go sample(message)
	//fmt.Println(<-message)
	//fmt.Println(<-message)
	//fmt.Println(<-message)
	//fmt.Println(<-message)
	// out put
	//Hello 1
	//Hello 2
	//Hello 3
	//Hello 4

	// use the chanel
	for msg := range message { // Note: no index for channel with range
		fmt.Println(msg)
	}
}
