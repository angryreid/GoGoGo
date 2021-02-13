package interface_test

import (
	"testing"
)

type Programmer interface {
	writeHelloWorld() string
}

type GoProgrammer struct {
}

// duck type
func (p *GoProgrammer) writeHelloWorld() string {
	return "fmt.Sprint(\"Hello World.\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.writeHelloWorld())
}
