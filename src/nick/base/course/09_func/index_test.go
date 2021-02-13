package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type MyIntFn func(op int) int

func returnMutilple2() (int, int) {
	return rand.Intn(10), rand.Intn(10)
}

func timeSpent(inner MyIntFn) MyIntFn {
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

// changable paramaters length
func getSum(ops ...int) int {
	ret := 0
	for _, v := range ops {
		ret += v
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(getSum(10, 20))
	t.Log(getSum(10, 20, 30))
}

func clear() {
	fmt.Print("Clear resources")
}

// defer func, fianlly execute
func TestDefer(t *testing.T) {
	t.Log("Started")

	defer clear()

	t.Log("Ongoing")

	panic("Fatal error") // urgent issue, but defer will execute, next code will be blocked
}
