package acacheimpl

import (
	"errors"
	"github.com/akyoto/cache"
	"time"
)

type AcacheImpl struct {
	Cache *cache.Cache
}

func NewCacheByACache(duration time.Duration)*AcacheImpl{
	gc := cache.New(duration)
	return &AcacheImpl{Cache:gc}
}

func (self *AcacheImpl) Get(key interface{})(interface{},error){
	val,ok := self.Cache.Get(key)
	if ok {
		return val,nil
	}
	return nil, errors.New("not found")
}

func (self *AcacheImpl) Set(key, value interface{})error{
	self.Cache.Set(key, value, 0)
	return nil
}

func (self *AcacheImpl) SetEx(key, value interface{}, expiration time.Duration)error{
	self.Cache.Set(key, value, expiration)
	return nil
}

func (self *AcacheImpl) Del(key string)error{
	self.Cache.Delete(key)
	return nil
}
