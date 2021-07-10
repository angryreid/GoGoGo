package tableTest

import (
	"math"
	"testing"
)

func add(a, b int32) int32 {
	return a + b
}

func TestTableData(t *testing.T) {
	tests := []struct {
		a, b, c int32
	}{
		{1, 2, 3},
		{2, 2, 4},
		{0, 0, 0},
		{math.MaxInt32, 1, math.MinInt32},
		{2, 3, 5},
	}
	for _, test := range tests {
		if calc := add(test.a, test.b); calc != test.c {
			// log.Printf("Test Failed At: %v", test)
			// return
			t.Errorf("add(%d, %d); get %d; expected %d", test.a, test.b, calc, test.c)

		}
	}
	//log.Println("All Test Cases passed")
	// t.Log("All Test Cases passed")
}
