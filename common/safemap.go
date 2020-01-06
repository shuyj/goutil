package common

import "sync"

type SMap struct {
	sync.RWMutex
	v map[string]interface{}
}

func NewSMap(cap int)*SMap{
	return &SMap{v:make(map[string]interface{}, cap)}
}

func (self *SMap) Put(key string, value interface{}) {
	self.Lock()
	defer self.Unlock()
	self.v[key] = value
}

func (self *SMap) Get(key string) (interface{},bool) {
	self.RLock()
	defer self.RUnlock()
	val,ok := self.v[key]
	return val,ok
}

func (self *SMap) Delete(key string) {
	self.Lock()
	defer self.Unlock()
	delete(self.v, key)
}

func (self *SMap) Size()int{
	self.RLock()
	defer self.RUnlock()
	return len(self.v)
}

func (self *SMap) Exists(key string) bool {
	self.RLock()
	defer self.RUnlock()
	_, ok := self.v[key]
	return ok
}
