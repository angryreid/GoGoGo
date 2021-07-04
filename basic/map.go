package main

import "fmt"

func main() {
	var m map[string]int // m === nil -> true
	m2 := make(map[string]string) // Zero Value
	fmt.Println(m, m2, m == nil)
	m3 := map[string]string {
		"name": "Nick",
		"age": "18",
		"skills": "JS TS GO Node HTML CSS Vue Angular React",
		"hobby": "guitar",
	}
	for k, v := range m3 {
		fmt.Println(k, v)
	}

	if name, ok := m3["name"]; ok {
		fmt.Println("The value is ", name)
	} else {
		fmt.Println("The key doesn't exist yet.")
	}

	delete(m3, "hobby")

	if hobby, ok := m3["hobby"]; ok {
		fmt.Println("The value is ", hobby)
	} else {
		fmt.Println("The key doesn't exist yet.")
	}
}
