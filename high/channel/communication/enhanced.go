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
		go func() { end <- true }()
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

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// waiting for all works finish
	for _, worker := range workers {
		<-worker.end
		<-worker.end
	}
}

func main() {
	chanDemo()
}
