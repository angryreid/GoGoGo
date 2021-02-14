package groutine_test

import (
	"fmt"
	"testing"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			// output is no order

		}(i)

		// go func() {
		// 	fmt.Println(i)
		// 	// all is 10
		// }()
	}

}
