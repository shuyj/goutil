package cache

import (
	"fmt"
	"github.com/goutil/common"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	Init()

	if err := Set("hello", "world"); err != nil {
		t.Error(err.Error())
	}
}

func TestGet(t *testing.T) {
	Init()

	val, err := Get("hello")
	if err != nil {
		t.Error(err.Error())
	}

	if val == nil {
		t.Error("val is nil")
	}
}

func TestAsyncNBlockCache(t *testing.T){
	Init()
	SetLimitSize(95)

	for i:=0; i<100; i++ {
		SetEx(fmt.Sprintf("key_%d", i), i, time.Duration(i)*time.Millisecond)
	}

	for i:=0; i<100; i++ {
		key := fmt.Sprintf("key_%d", i)
		val, rtime, err := GetWithTime(key)

		if err != nil {
			fmt.Printf("Get error = %v\n", err)
			continue
		}
		//fmt.Printf("remain timme = %v", rtime)
		if rtime <= 10*time.Millisecond {
			// expired
			// reload asynchronous
			common.Async(func() {
				// get data with key
				fmt.Printf("reload key=%s\n", key)
				SetEx(key, i, 10*time.Second)
			})
		}
		// return cached data whether or not expired
		val=val
	}

	common.WaitAndStop()

}




