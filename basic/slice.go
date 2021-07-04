package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("Slice: %vLen: %d, Cap: %d \n", s, len(s), cap(s))
}

func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Println("arr[2:6]", arr[2:6])
	fmt.Println("arr[2:6]", arr[:])
	// Once the length over the Capacity, The slice will refer to a new Slice with greater capacity
	// Creating Slice
	var s []int
	for i := 0; i < 100; i++ {
		//printSlice(s)
		s = append(s,  2 * i + 1)
	}

	s1 := []int{2,3,4,5}
	printSlice(s1)

	s2 := make([]int, 10, 16)
	printSlice(s2)

	// Coping Slice
	copy(s2, s1)
	printSlice(s2)

	// Deleting Slice
	s2 = append(s2[:2], s2[3:]...)
	printSlice(s2)

	// Popping from front
	front := s2[0]
	s2 = s2[1:]
	tail := s2[len(s2) - 1]
	s2 = s2[:len(s2) -1]
	// Popping from tail
	fmt.Println(front, tail)
	printSlice(s2)
}
