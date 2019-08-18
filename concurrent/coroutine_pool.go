package concurrent

import (
	"errors"
	"time"
)

// 对象池
type ObjPoll struct {
	bufChan chan interface{}
}

// 创建对象池
func NewObjPoll(obj interface{}, num int) *ObjPoll {
	poll := ObjPoll{}
	poll.bufChan = make(chan interface{}, num)
	for i := 0; i < num; i++ {
		poll.bufChan <- obj
	}

	return &poll
}

// 获取对象
func (p *ObjPoll) GetObj(timeout time.Duration) (interface{}, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

// 放回对象
func (p *ObjPoll) Release(obj interface{}) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}
