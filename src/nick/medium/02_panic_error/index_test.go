package test_panic

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicInit(t *testing.T) {
	defer func() {
		fmt.Println("Hello")
	}()

	fmt.Println("starting")
	panic(errors.New("somthing wrong")) // show call stack, execute defer

	// os.Exit(-1)
	fmt.Println("end")
}

func TestRecoverInit(t *testing.T) {
	defer func() {
		if err := recover(); err != nil { // recover handle panic error, the process will keep executing...
			fmt.Println("recover from", err)
		}
	}()

	fmt.Println("starting")
	panic(errors.New("somthing wrong")) // show call stack, execute defer

	// os.Exit(-1)
	fmt.Println("end")
}
