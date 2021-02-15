package obj_pool

import (
	"errors"
	"time"
)

type ResusableObj struct {
}

type ObjPool struct {
	bufChan chan *ResusableObj
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ResusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

// PushObj ...
func (p *ObjPool) PushObj(obj *ResusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("channel size overflow.")
	}
}

func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}

	objPool.bufChan = make(chan *ResusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ResusableObj{}
	}
	return &objPool
}
