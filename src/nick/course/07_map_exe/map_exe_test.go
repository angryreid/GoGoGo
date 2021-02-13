package map_test

import "testing"

func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}

	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op - 1 }

	t.Log(m[1](2))
	t.Log(m[2](2))
	t.Log(m[3](1))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}

	mySet[1] = true

	// delete(mySet, 1) // (set, key)
	n := 1
	if mySet[n] {
		t.Logf("%d is existing.", n)
	} else {
		t.Logf("%d is not existing.", n)
	}

	t.Logf("length is %d ", len(mySet))

}
