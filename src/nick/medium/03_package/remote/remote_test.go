package remote_test

import (
	"testing"

	cm "github.com/easierway/concurrent_map"
)

func TestRemoteInit(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("year"), 2021)
	t.Log(m.Get(cm.StrKey("year")))
}
