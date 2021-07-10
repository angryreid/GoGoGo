package main

import (
	"fmt"
	"runtime"
	"time"
)

func testGoRoutine() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				//fmt.Printf("Hello from goroutine %d\n", i)
				arr[i]++
				runtime.Gosched() // release control
			}
		}(i)
	}
	// To avoid main func close the space
	time.Sleep(time.Millisecond)

	fmt.Println(arr)
}

func main() {
	testGoRoutine()
}
