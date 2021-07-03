package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello")
	const filename = "testing.txt"
	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Printf("%s\n", contents)
	} else {
		fmt.Println(err)
	}
}
