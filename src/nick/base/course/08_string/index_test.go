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

func TestStringRune(t *testing.T) {
	s := "我最喜欢源源了"

	for _, c := range s {
		t.Logf("%[1]c %[1]d", c) // use first parameters too. but %[1]c is char, %[1]d is code
	}
}
