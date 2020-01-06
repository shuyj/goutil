package common

import (
	"encoding/json"
	"log"
	"testing"
	"time"
)

type DataInfo struct {
	IndexType string      `json:"index_type"`
	Timestamp int64       `json:"time"`
	Hostip    string      `json:"hostip"`
	Body      interface{} `json:"body"`
}

type Inst struct {
}

func (self *Inst) OnElement(i interface{}){
	jsondata, _ := json.Marshal(i)
	//nw,err := self.Writer.Write(jsondata)
	log.Printf("Datalog write = %v", string(jsondata))
}

func TestNewSyncToAsyncInterface(t *testing.T) {
	channelMaxLen := 1000
	coroutineNum := 2
	inst := &Inst{}
	Queue := NewSyncToAsyncInterface(channelMaxLen, coroutineNum, inst)
	err := Queue.Start()
	if err != nil {
		log.Printf("Datalog queue start err = %v", err)
		return
	}

	info := &DataInfo{
		IndexType: "test",
		Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
		Body:   "{\"hello\":\"world\"}",
		Hostip:    "0.0.0.0",
	}
	err = Queue.Send(info)
	if err != nil {
		log.Printf("Datalog in queue  err = %v", err)
	}
}


func TestNewSyncToAsyncWithChan(t *testing.T) {
	channelMaxLen := 1000
	coroutineNum := 2
	playerCount := NewSyncToAsyncWithChan(channelMaxLen, coroutineNum, func(i interface{}) {
		log.Printf(" Count Debug %+v", i)
	})
	err := playerCount.Start()
	if err != nil {
		log.Printf(" NewSyncToAsyncWithChan Start Error : %+v", err)
	}

	playerCount.Send(nil)
}

