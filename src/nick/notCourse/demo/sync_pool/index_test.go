package sync_pool_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creat a new object.")
			return 100
		},
	}

	v := pool.Get().(int)

	fmt.Println(v)

	pool.Put(33)

	v1, _ := pool.Get().(int)

	fmt.Println(v1)
}
