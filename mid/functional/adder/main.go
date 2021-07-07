package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type IAdder func(int) (int, IAdder)

func adder2 (base int) IAdder {
	return func(v int) (int, IAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("Index: %d Sum: %d \n", i, a(i))
	}

	var sum int
	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		sum, a2 = a2(i)
		fmt.Printf("Index: %d Sum: %d \n", i, sum)
	}
}
