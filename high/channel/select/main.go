package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func testSelect() {
	var c1, c2 = generator(), generator()
	ta := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		select {
		case n := <-c1:
			fmt.Println("Received from c1:", n)
		case n := <-c2:
			fmt.Println("Received from c2:", n)
		case <-tick:
			fmt.Println("Time Tick")
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Time out")
		case <-ta:
			fmt.Println("Bye")
			return
		}
	}
}

func main() {
	testSelect()
}
