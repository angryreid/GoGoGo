package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct{}

func (p *Pet) speak() {
	fmt.Print("...")
}

func (p *Pet) speakTo(host string) {
	p.speak()
	fmt.Println(" ", host)
}

type Dog struct{}

func (d *Dog) speak() {
	fmt.Print("Wang wang wang~ ")
}

func (d *Dog) speakTo(host string) {
	d.speak()
	fmt.Println(" ", host)
}

func TestDogExtension(t *testing.T) {
	dog := new(Dog)

	dog.speakTo("Nick")
}
