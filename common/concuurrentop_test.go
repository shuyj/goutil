package common

import (
	"log"
	"testing"
	"time"
)

const defaultConfUpdateInterval = time.Second * 1
var xxxconfig = NewConcurrentOp([]string{}, defaultConfUpdateInterval, realGetCustomConfig)

func realGetCustomConfig(key string)(interface{},error){
	//value, err := redisclient.GetSlave().SMembers(key)
	var err error
	value := []string{"hello", "world"}
	log.Printf("real read config %s = %+v", key, value)
	if err != nil {
		return []string{}, err
	}
	return value,nil
}


func TestConcurrentOp_ConcurrentGet(t *testing.T) {
	val, err := xxxconfig.ConcurrentGet("xxx.config.custom.config")
	if err != nil {
		return
	}
	rval, ok := val.([]string)
	if !ok {
		return
	}

	log.Printf("query result = %v", rval)
}


func BenchmarkConcurrentOp_ConcurrentGet(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			xxxconfig.ConcurrentGet("xxx.config.custom.config")
		}
	})
}




