package extension_test

import (
	"fmt"
	"testing"
)

type Cat struct {
	Pet
}

func (c *Cat) speak() {
	fmt.Print("Miao miao miao~")
}

func TestCatExtension(t *testing.T) {
	cat := new(Cat)

	cat.speakTo("Nick") // ... Nick
}
