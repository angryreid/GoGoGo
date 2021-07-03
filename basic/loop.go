package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	res := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		res = strconv.Itoa(lsb) + res
	}
	return res
}

func printFile(filename string) {
	if file, err := os.Open(filename); err != nil {
		panic(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}

}

func eval(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a -b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Not support" + op)
	}
}

func div (a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	fmt.Println("loop")
	fmt.Println(
		convertToBin(7),
		convertToBin(15),
		convertToBin(31))
	printFile("testing.txt")
	fmt.Println(eval(1,2, "+"))
	q, r := div(17,2)
	fmt.Println(q, r)
}
