package common

import (
	"sync/atomic"
	"time"
)

type OpenLock struct {
	Version uint64
}

func NewOpenLock()*OpenLock{
	return &OpenLock{Version:0}
}

func (self *OpenLock) Acquire(ov uint64) bool {
	if self.Version != ov {
		return false
	}
	if atomic.CompareAndSwapUint64(&self.Version, ov, self.Version+1) {
		return true
	}
	return false
}


func AcquireForOneNoBlock(lck *OpenLock, cond func()bool, exec func()){
	oversion := lck.Version
	if cond() {
		if lck.Acquire(oversion) {
			Async(exec)
		}
	}
}

const(
	LockStateLocked = 1
	LockStateUnlock = 0
)

type NoBlockLock struct {
	State uint64
}

func NewNoBlockLock()*NoBlockLock{
	return &NoBlockLock{State:LockStateUnlock}
}

func (self *NoBlockLock) Lock() bool {
	if atomic.CompareAndSwapUint64(&self.State, LockStateUnlock, LockStateLocked) {
		return true
	}
	return false
}

func (self *NoBlockLock) Unlock(){
	atomic.CompareAndSwapUint64(&self.State, LockStateLocked, LockStateUnlock)
}

func AcquireForOneNoBlockLock(lck *NoBlockLock, cond func()bool, exec func(), duration time.Duration){
	if cond() {
		if lck.Lock() {
			Async(exec)
			time.AfterFunc(duration, func() {
				lck.Unlock()
			})
		}
	}
}

type NoBlockSem struct {
	count int64
}

func NewNoBlockSem(defnum int64)*NoBlockSem{
	return &NoBlockSem{count:defnum}
}

func (self *NoBlockSem) Acquire()bool{
	old := self.count
	if old <= 0 {
		return false
	}
	if atomic.CompareAndSwapInt64(&self.count, old, old-1) {
		return true
	}
	return false
}

func (self *NoBlockSem) Release(){
	atomic.AddInt64(&self.count, 1)
}


