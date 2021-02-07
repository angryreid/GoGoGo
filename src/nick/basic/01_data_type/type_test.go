package data

import "testing"

type MyInt int64

func TestDataType(t *testing.T) {
	var a int64 = 1
	var b int32

	var c MyInt

	// b = a // wrong
	b = int32(a)

	// c = a // Error: alias can't change
	c = MyInt(a)

	t.Log(a, b, c)

}

func TestPointer(t *testing.T) {
	a := 1

	aPtr := &a

	// aPtr = aPtr + 1 // Error: not support

	t.Log(a, aPtr)

	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string

	t.Log("*" + s + "*") // **
	t.Log(len(s))        // 0
}
