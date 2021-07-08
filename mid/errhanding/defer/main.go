package main

import (
	"bufio"
	"fmt"
	"go-learing/mid/functional/fib"
	"os"
)

func deferTry() {
	defer fmt.Println(1) // FILO Stack
	defer fmt.Println(2)
	fmt.Println(3)
	panic("Defer will call before panic")
}

func OpenErrorFile(filename string) {
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		//panic(err)
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
	}
	defer file.Close()
}

func WriteFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	// deferTry()
	const filename = "emma.txt"
	// WriteFile(filename)
	// When to use defer ?
	//1. Open / Close
	//2. Lock / Unlock
	//3. PrintHeader / PrintFooter

	OpenErrorFile(filename)
}
