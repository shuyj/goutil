package jsonutil

import (
	"encoding/json"
	"log"
)

type (
	Writer interface {
		AddField(name string, val interface{}) Writer
		RemoveField(name string) Writer
		HasField(name string) bool
		GetField(name string) interface{}
		BuildJson() (string,error)
	}

	writerImpl struct {
		fields map[string]interface{}
	}
)

func NewJsonWriter() Writer {
	return &writerImpl{
		fields: map[string]interface{}{},
	}
}

func (self *writerImpl) AddField(name string, val interface{}) Writer {
	self.fields[name] = val
	return self
}

func (self *writerImpl) RemoveField(name string) Writer {
	delete(self.fields, name)

	return self
}

func (self *writerImpl) HasField(name string) bool {
	_, ok := self.fields[name]
	return ok
}

func (self *writerImpl) GetField(name string) interface{} {
	if !self.HasField(name) {
		return nil
	}
	return self.fields[name]
}

func (self *writerImpl) BuildJson() (string,error){
	for k, v := range self.fields {
		log.Printf("BuildJson %s=%T|%v", k, v, v)
	}

	bytes, err := json.Marshal(self.fields)

	if err != nil {
		return "",err
	}
	return string(bytes),nil
}

