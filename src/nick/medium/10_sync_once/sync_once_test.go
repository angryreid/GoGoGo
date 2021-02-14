package sync_once

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func getSingleton() *Singleton {
	once.Do(func() {
		fmt.Println("Creat instance.")

		singleInstance = new(Singleton)
	})

	return singleInstance
}

func TestGetSingleton(t *testing.T) {
	var wg sync.WaitGroup // for wait all groutins running done
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			singleIns := getSingleton()
			fmt.Printf("%x\n", unsafe.Pointer(singleIns))
			wg.Done()
		}()
	}
	wg.Wait()
}
