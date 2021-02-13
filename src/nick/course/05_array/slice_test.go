package array_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s []int
	t.Log(len(s), cap(s))
	s = append(s, 1)
	t.Log(len(s), cap(s))

	s1 := make([]int, 3, 5)

	t.Log(len(s1), cap(s1))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}

	for i := 0; i < 10; i++ {
		s = append(s, i)

		t.Log(len(s), cap(s), i)
	}
}

func TestSliceShareMemory(t *testing.T) {
	year := []string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sep",
		"Oct",
		"Nov",
		"Dec",
	}

	session2 := year[3:6]

	t.Log(session2, len(session2), cap(session2))

	summer := year[5:8]

	t.Log(summer, len(summer), cap(summer))

	summer[0] = "Unknow"

	t.Log(session2)

}
