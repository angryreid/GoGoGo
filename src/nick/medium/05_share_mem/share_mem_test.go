package share_mem

import (
	"sync"
	"testing"
	"time"
)

func TestSharedMemory(t *testing.T) {
	counter := 0
	for i := 0; i < 100; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 1)
	t.Logf("counter= %d", counter)
}

func TestSharedMemorySafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 100; i++ {

		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(time.Second * 1) // make sure all groutine running donw
	t.Logf("counter= %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	t.Logf("counter= %d", counter)
}
