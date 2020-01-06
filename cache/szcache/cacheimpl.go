package szcache

import (
	"container/list"
	"errors"
	"fmt"
	"github.com/goutil/common"
	"log"
	"sync"
	"time"
)

type Cache struct {
	items *common.SMap
	size  uint64
	sizeLimit uint64
	close chan struct{}
	//queue *common.FIFOQueue
	queue *list.List
	lmux  sync.RWMutex
}

type item struct {
	key 	string
	data    interface{}
	expires int64
}

func NewCacheBySize(maxSize uint64) *Cache {
	cache := &Cache{
		close: make(chan struct{}),
		sizeLimit:maxSize,
		size:0,
		//queue: common.NewFIFOQueue(),
		queue:list.New(),
		items:common.NewSMap(int(maxSize/2)),
	}

	return cache
}

func (cache *Cache) SetLimitSize(limit uint64){
	cache.sizeLimit = limit
	log.Printf("szCache update limit size = %d", limit)
}

func (cache *Cache) Get(key string) (interface{}, error) {
	obj, exists := cache.items.Get(key)

	if !exists {
		return nil, errors.New("not exists")
	}

	iem,ok := obj.(*list.Element)
	if !ok {
		return nil, errors.New("type not list.Element")
	}
	item,ok := iem.Value.(*item)
	if !ok {
		return nil, errors.New("type not item")
	}

	if item.expires > 0 && time.Now().UnixNano() > item.expires {
		return nil, errors.New("expired")
	}

	return item.data, nil
}

func (cache *Cache) GetWithTime(key string) (interface{}, time.Duration, error) {
	obj, exists := cache.items.Get(key)

	if !exists {
		return nil, 0, errors.New("not exists")
	}

	iem,ok := obj.(*list.Element)
	if !ok {
		return nil, 0, errors.New("type not list.Element")
	}
	item,ok := iem.Value.(*item)
	if !ok {
		return nil, 0, errors.New("type not item")
	}
	now := time.Now().UnixNano()

	if item.expires > 0 && now > item.expires {
		return item.data, 0, nil
	}

	return item.data, time.Duration(item.expires-now), nil
}

func (cache *Cache) Set(key string, value interface{}, duration time.Duration) error {
	var expires int64
	if duration > 0 {
		expires = time.Now().Add(duration).UnixNano()
	}
	var upsize = false
	if cache.deleteQueue(key) != nil {
		upsize = true
	}
	cache.lmux.Lock()
	elem := cache.queue.PushBack(&item{
		key: key,
		data:    value,
		expires: expires,
	})
	cache.lmux.Unlock()
	cache.items.Put(key, elem)
	if upsize {
		cache.size = uint64(cache.items.Size())
	}

	for cache.size > cache.sizeLimit {
		fmt.Printf("FULL limit=%d-%d-%d\n", cache.sizeLimit, cache.size, cache.items.Size())
		cache.lmux.RLock()
		em := cache.queue.Front()
		cache.lmux.RUnlock()
		cache.deleteElement(em)
		//item,ok := em.Value.(*item)
		//if ok {
		//	cache.Delete(item.key)
		//}
	}
	return nil
}

func (cache *Cache) deleteElement(elem *list.Element) error {
	cache.lmux.Lock()
	cache.queue.Remove(elem)
	cache.lmux.Unlock()
	item,ok := elem.Value.(*item)
	if ok {
		cache.items.Delete(item.key)
		cache.size = uint64(cache.items.Size())
	}
	return nil

}

func (cache *Cache) deleteQueue(key string) error {
	elem, ok := cache.items.Get(key)
	if ok {
		cache.lmux.Lock()
		cache.queue.Remove(elem.(*list.Element))
		cache.lmux.Unlock()
		return nil
	}
	return errors.New("not exists")
}

func (cache *Cache) Delete(key string) error {
	elem, ok := cache.items.Get(key)
	if ok {
		cache.lmux.Lock()
		cache.queue.Remove(elem.(*list.Element))
		cache.lmux.Unlock()
		cache.items.Delete(key)
		cache.size = uint64(cache.items.Size())
		return nil
	}
	return errors.New("not exists")
}

func (cache *Cache) Close() {
	cache.close <- struct{}{}
	cache.items = nil
}

func (cache *Cache) Count() uint64{
	return uint64(cache.size)
}

