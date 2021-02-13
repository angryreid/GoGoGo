package empty_interface

import (
	"fmt"
	"testing"
)

func doSomething(p interface{}) {
	if i, ok := p.(int); ok {
		fmt.Println("Integer ", i)
		return
	}

	if i, ok := p.(string); ok {
		fmt.Println("String ", i)
		return
	}

	fmt.Println("Unknown Type")
}

func doSomethingEnhance(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer ", v)
	case string:
		fmt.Println("String ", v)
	default:
		fmt.Println("Unknown Type")
	}

}

func TestEmptyInterfaceAssertion(t *testing.T) {
	doSomethingEnhance(10)
	doSomethingEnhance("10")
	doSomethingEnhance(1.2)
}
