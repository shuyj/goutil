package gocacheimpl

import (
	"errors"
	"github.com/patrickmn/go-cache"
	"time"
)

type GoCacheImpl struct {
	Cache *cache.Cache
}

func NewCacheByGoCache(duration time.Duration)*GoCacheImpl{
	gc := cache.New(5*time.Minute, 10*time.Minute)
	return &GoCacheImpl{Cache:gc}
}

func (self *GoCacheImpl) Get(key string)(interface{},error){
	val, ok := self.Cache.Get(key)
	if ok {
		return val,nil
	}
	return nil, errors.New("not found")
}

func (self *GoCacheImpl) Del(key string)error{
	self.Cache.Delete(key)
	return nil
}

func (self *GoCacheImpl) Set(key string, value interface{})error{
	self.Cache.Set(key, value, cache.NoExpiration)
	return nil
}

func (self *GoCacheImpl) SetEx(key string, value interface{}, expiration time.Duration)error{
	self.Cache.Set(key, value, expiration)
	return nil
}

func (self *GoCacheImpl) Count()int{
	return self.Cache.ItemCount()
}
