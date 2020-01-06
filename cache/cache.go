package cache

import (
	"github.com/goutil/cache/szcache"
	"sync"
	"time"
)

var (
	singleInstance sync.Once
	gocache        *szcache.Cache
)

func Init() {
	singleInstance.Do(func() {
		gocache = szcache.NewCacheBySize(10000)
	})
}

func Set(key string, value interface{}) error {
	return gocache.Set(key, value, time.Hour)
}

func SetEx(key string, value interface{}, expire time.Duration) error {
	return gocache.Set(key, value, expire)
}

func Get(key string) (interface{}, error) {
	return gocache.Get(key)
}

func GetWithTime(key string)(interface{}, time.Duration, error){
	return gocache.GetWithTime(key)
}

func Count() uint64 {
	return gocache.Count()
}

func Del(key string)error{
	return gocache.Delete(key)
}

func SetLimitSize(limit uint64){
	gocache.SetLimitSize(limit)
}
