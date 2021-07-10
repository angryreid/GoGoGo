package tableTest

import "testing"

func BenchMarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if calc := add(1, 2); calc != 3 {
			b.Error("Error")
		}
	}
}
