package main

import (
	"fmt"
	"time"
)

func nextStringChan(message chan string) {
	for i := 0; i < 10; i++ {
		message <- "message:" + string(('A' + i))
		time.Sleep(2 * time.Second)
	}
}

func nextIntChan(number chan int) {
	for i := 0; i < 5; i++ {
		number <- i
		time.Sleep(time.Second)
	}
}

func main() {
	number := make(chan int)
	message := make(chan string)

	go nextIntChan(number)
	go nextStringChan(message)

	//go func() {
	//	for {
	//		select {
	//		case num, _ := <-number:
	//			fmt.Println(num)
	//		case msg, _ := <-message:
	//			fmt.Println(msg)
	//		}
	//	}
	//}()
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error coming:", err)
		} else {
			panic(r)
		}
	}()

	for {
		select {
		case num, _ := <-number:
			fmt.Println(num)
		case msg, _ := <-message:
			fmt.Println(msg)
		}
	}

	// fatal error: all goroutines are asleep - deadlock!

}
