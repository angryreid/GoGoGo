package polym_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	writeHelloWorld() Code
}

type GoProgrammer struct{}

func (p *GoProgrammer) writeHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct{}

func (p *JavaProgrammer) writeHelloWorld() Code {
	return "System.out.println(\"Hello World\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T, %v\n", p, p.writeHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goPro := new(GoProgrammer)
	javaPro := new(JavaProgrammer)
	writeFirstProgram(goPro)
	writeFirstProgram(javaPro)
}
