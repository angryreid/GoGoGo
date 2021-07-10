package main

import (
	"fmt"
	"time"
)

func bufferedChannel() {
	c := make(chan int, 3)
	worker(0, c)
	c <- 'a' + 1
	c <- 'a' + 2
	c <- 'a' + 3
	//c <- 4 // deadlock
	close(c) // only sender can close
	// Communication Sequential Process CSP
	time.Sleep(time.Millisecond)
}

func worker(id int, c <-chan int) {
	go func() {
		for {
			if n, ok := <-c; !ok {
				break
			} else {
				fmt.Printf("Worker ID: %d has received data %c \n",
					id, n)
			}
		}
	}()
}

func main() {
	bufferedChannel()
}
