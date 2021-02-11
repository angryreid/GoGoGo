package string_test

import "testing"

func TestStringInit(t *testing.T) {
	var s string

	t.Log("*" + s + "*") // **

	s = "新年快乐"

	t.Log(len(s)) // byte numbers

	// s = "\xE4\xB8\xA5"

	// t.Log(s)

	c := []rune(s)

	t.Log(len(c))

	t.Logf("unicode %x", c[0])
	t.Logf("utf-8 %x", s)
}
