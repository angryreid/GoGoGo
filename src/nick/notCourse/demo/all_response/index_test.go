package all_response

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The searching result is coming from %d", id)
}

func AllResponse() string {
	numOfRunner := 10

	// using buffer channel to avoid out of memory
	ch := make(chan string, numOfRunner)
	// ch := make(chan string)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	var finalRet string
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}

	return finalRet
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log("\n" + AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("End:", runtime.NumGoroutine())

}
