package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	// use : to type asset
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}

	fmt.Println("All: ", b)

	t.Log("Fib ", b)
}

func TestExchangeTwo(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a

	t.Log(a, b)
}
