package common

import (
	"errors"
	"time"
)

type refreshFunc func(string) (interface{}, error)

type ConcurrentOp struct {
	cacheValue     interface{}
	updateInterval time.Duration
	lastUpdateTime time.Time
	openLock       *OpenLock
	updateFunc     refreshFunc
}

func NewConcurrentOp(defValue interface{}, updateInterval time.Duration, updatefunc refreshFunc) *ConcurrentOp {
	return &ConcurrentOp{cacheValue: defValue, updateInterval: updateInterval, lastUpdateTime: time.Now().Add(-updateInterval), openLock: NewOpenLock(), updateFunc: updatefunc}
}

func (self *ConcurrentOp) ConcurrentGet(key string)(retcode interface{},reterr error){
	defer func() {
		if e := recover(); e != nil {
			retcode = self.cacheValue
			reterr = e.(error)
		}
	}()

	AcquireForOneNoBlock(self.openLock,
		func()bool {
			now := time.Now()
			if self.lastUpdateTime.Add(self.updateInterval).Before(now) {
				self.lastUpdateTime = now
				return true
			}
			return false
		},
		func(){
			var err error
			var val interface{}
			val, err = self.Get(key)
			if err == nil {
				self.cacheValue = val
			}
		})
	return self.cacheValue, nil
}

func (self *ConcurrentOp) Get(key string)(interface{},error){
	if self.updateFunc != nil {
		return self.updateFunc(key)
	}
	return 0, errors.New("must be set update Function")
}



