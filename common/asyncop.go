package common

import (
	"errors"
	"log"
	"runtime"
	"sync"
)

type ExecFunction func(interface{})

type ExecInterface interface{
	OnElement(interface{})
}

type SyncToAsyncWithChan struct {
	channMsg     chan interface{}
	startFlag    bool
	coroutineNum int
	channelMaxLen int
	execFunction ExecFunction
	execInterface ExecInterface
	wg 			sync.WaitGroup
}

func NewSyncToAsyncWithChan(channelMaxLen, coroutineNum int, execfunc ExecFunction) *SyncToAsyncWithChan {
	return &SyncToAsyncWithChan{
		channMsg:     make(chan interface{}, channelMaxLen),
		startFlag:    false,
		channelMaxLen: channelMaxLen*95/100, // limit 95 precent
		coroutineNum: coroutineNum,
		execFunction: execfunc,
	}
}

func NewSyncToAsyncInterface(channelMaxLen, coroutineNum int, execif ExecInterface) *SyncToAsyncWithChan {
	return &SyncToAsyncWithChan{
		channMsg:     make(chan interface{}, channelMaxLen),
		startFlag:    false,
		channelMaxLen: channelMaxLen*95/100, // limit 95 precent
		coroutineNum: coroutineNum,
		execInterface: execif,
	}
}

func (self *SyncToAsyncWithChan) Start() error {
	if self.startFlag {
		return nil
	}
	if self.execFunction == nil && self.execInterface == nil {
		return errors.New("SyncToAsyncWithChan exec function or interface must be set")
	}

	var coroutineNum int
	if self.coroutineNum < 1 {
		coroutineNum = runtime.NumCPU()
	} else {
		coroutineNum = self.coroutineNum
	}
	for index := 0; index < coroutineNum; index++ {
		self.wg.Add(1)
		go self.onWork()
	}

	log.Printf("SyncToAsyncWithChan start ok, coroutine number:%d channel max len:%d", coroutineNum, self.channelMaxLen)

	self.startFlag = true
	return nil
}

func (self *SyncToAsyncWithChan) Stop(waitForAll bool) {
	if !self.startFlag {
		return
	}

	self.startFlag = false
	close(self.channMsg)
	if waitForAll {
		self.wg.Wait()
	}
}

func (self *SyncToAsyncWithChan) Send(val interface{}) error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("SyncToAsyncWithChan writer panic:%v", r)
			return
		}
	}()

	if !self.startFlag {
		log.Print("SyncToAsyncWithChan not init")
		return errors.New("SyncToAsyncWithChan not init")
	}

	if len(self.channMsg) > self.channelMaxLen {
		log.Printf("message channel is full, limit len:%d", self.channelMaxLen)
		return errors.New("message channel is full")
	}
	//log.Debug("SyncToAsyncWithChan msg = ", string(value))
	self.channMsg <- val
	return nil
}

func (self *SyncToAsyncWithChan) onWork() {
	for {
		val, ok := <-self.channMsg
		if !ok {
			log.Printf("message channel is close")
			break
		}
		if self.execFunction != nil {
			self.execFunction(val)
		}else if self.execInterface != nil{
			self.execInterface.OnElement(val)
		}
	}
	self.wg.Done()
}
