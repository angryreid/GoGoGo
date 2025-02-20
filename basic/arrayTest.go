package main

import "fmt"

func printArray(arr [5]int)  {
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{1,2,3,4,5}
	var grid [2][3]int

	fmt.Println(arr1, arr2, arr3, grid)
	printArray(arr1) // array is value coping
}
