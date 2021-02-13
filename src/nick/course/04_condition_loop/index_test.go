package loop

import "testing"

// loop
func TestWhileLoop(t *testing.T) {
	n := 0
	/* while (n < 5) */
	for n < 5 {
		t.Log(n)
		n++
	}
}

// if
func TestIfMultiSec(t *testing.T) {
	t.Log("If")
	if a := 1 == 1; a {
		t.Log("1 == 1")
	}

	// if v, err := someFunc(); err == nil {
	// 	//
	// } else {
	// 	//
	// }
}

func TestSwitch(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not in 0-3.")
		}
	}
}

func TestSwitchAsIf(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 != 0:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}
}
