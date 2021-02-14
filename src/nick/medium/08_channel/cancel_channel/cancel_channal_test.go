package cancel_channel

import (
	"fmt"
	"testing"
	"time"
)

func cancelOneChannel(cancelCh chan struct{}) {
	cancelCh <- struct{}{}
}
func cancelAllChannel(cancelCh chan struct{}) {
	close(cancelCh)
}

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func TestCalcel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled.")
		}(i, cancelChan)
	}
	// cancelOneChannel(cancelChan) // 4 concelled
	cancelAllChannel(cancelChan) // all cancelled
	time.Sleep(time.Second * 1)
}
