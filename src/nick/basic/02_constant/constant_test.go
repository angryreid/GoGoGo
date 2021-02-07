package constant

import "testing"

const (
	Monday = 1 + iota
	Tuesday
	Wednesday
)

const (
	Readable = 1 << iota
	Writable
	Excutable
)

func TestConstTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConsantBit(t *testing.T) {
	a := 7 // 0111
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Excutable == Excutable)
}
