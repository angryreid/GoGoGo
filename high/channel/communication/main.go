package main

import (
	"fmt"
)

type worker struct {
	in  chan int
	end chan bool
}

func doWorker(id int, c chan int, end chan bool) {
	for n := range c {
		fmt.Printf("Worker ID: %d has received data %c\n", id, n)
		end <- true
	}
}

func createWorker(id int) worker {
	w := worker{
		in:  make(chan int),
		end: make(chan bool),
	}
	go doWorker(id, w.in, w.end)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
		<-workers[i].end
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].end
	}
}

func main() {
	chanDemo()
}
