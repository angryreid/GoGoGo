package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int

	arr1 := [3]int{1, 2, 3}

	arr2 := [...]int{1, 2, 3, 4, 5}

	t.Log(arr, arr1, arr2)
}

func TestArrayTravel(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		t.Log(arr[i])
	}

	for idx, v := range arr {
		t.Log(idx, v)
	}

}

func TestArraySection(t *testing.T) {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	arr_sec := arr[2:]
	t.Log(arr_sec)
}
