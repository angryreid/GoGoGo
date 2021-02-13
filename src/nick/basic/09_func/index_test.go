package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMutilple2() (int, int) {
	return rand.Intn(10), rand.Intn(10)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())

		return ret
	}
}

func slowFn(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsFS := timeSpent(slowFn)
	t.Log(tsFS(10))
}

func TestMutilpleReturn2Value(t *testing.T) {
	// a, b := returnMutilple2()
	a, _ := returnMutilple2()
	// t.Log(a, _) // cannot use _ as a value or type (it's using for ignore variable)
	t.Log(a)
}
