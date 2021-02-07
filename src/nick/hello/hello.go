// entry package main !!!

package main

import (
	"fmt"
	"os"
)

// main entry !!!
// no return type.
// os.Exit to return status
// os.Args to access paramaters
func main() {

	if len(os.Args) > 1 {
		fmt.Println("Parameters is :", os.Args[1])
	}
	fmt.Println("hello, world!")
	os.Exit(1)
}
