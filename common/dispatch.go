package common

import (
	"log"
)

var dispatch *SyncToAsyncWithChan

func init(){
	dispatch = NewSyncToAsyncWithChan(2000, -1, func(i interface{}) {
		if f, ok := i.(func()); ok {
			f()
		}
	})
	err := dispatch.Start()
	if err != nil {
		log.Printf("dispatch create error=%s", err.Error())
	}
	log.Printf("Dispatch component init ")
}

func Async(f func())error{
	return dispatch.Send(f)
}

func WaitAndStop(){
	dispatch.Stop(true)
}




