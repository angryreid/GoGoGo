package main

import (
	"fmt"
)

func tryRecover(){
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error coming:", err)
		} else {
			panic(r)
		}
	}()
	//panic(errors.New("This is an error"))
	a,b := 1,0
	c := a/b
	fmt.Print(c)
}

func main() {
	tryRecover()
}
