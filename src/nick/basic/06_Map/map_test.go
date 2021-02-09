package map_test

import "testing"

func TestMapInit(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m["one"])
	m1 := map[string]int{}
	t.Log(m1)
	m2 := make(map[string]int, 10 /*Initial Capacity*/)
	t.Log(m2)

	t.Logf("m len=%d", len(m))
	t.Logf("m1 len=%d", len(m1))
	t.Logf("m2 len=%d", len(m2))
}

func TestAccessNotExistingKey(t *testing.T) {
	m := map[int]int{}

	t.Log(m[1])

	m[2] = 0

	t.Log(m[2])

	m[3] = 0

	fourthValue, isExist := m[4]

	t.Log(fourthValue, isExist)

	if v, ok := m[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("Key 3 is not existing.")
	}
}

func TestTravelMap(t *testing.T) {
	m := map[int]int{1: 1, 2: 2, 3: 3, 4: 5, 5: 5}
	for key, value := range m {
		t.Log(key, value)
	}
}
