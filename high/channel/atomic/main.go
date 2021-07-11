package main

import (
	"fmt"
	"time"
)

type AtomicInt int

func (a *AtomicInt) increment() {
	*a++
}

func (a *AtomicInt) get() int {
	return int(*a)
}

func main() {
	var a AtomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
