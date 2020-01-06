package common

import (
	"errors"
	"fmt"
	"sync"
)

type FIFOQueue struct {
	sync.RWMutex
	slice       []interface{}
}

func NewFIFOQueue() *FIFOQueue {
	ret := &FIFOQueue{}
	ret.slice = make([]interface{}, 0)
	return ret
}

func (st *FIFOQueue) Exists(value interface{})bool{
	st.RLock()
	defer st.RUnlock()
	for _,v := range st.slice{ // O(N) slower
		if v == value{
			return true
		}
	}
	return false
}

func (st *FIFOQueue) Enqueue(value interface{}) error {
	st.Lock()
	defer st.Unlock()
	st.slice = append(st.slice, value)
	return nil
}

func (st *FIFOQueue) Dequeue() (interface{}, error) {
	st.Lock()
	defer st.Unlock()

	len := len(st.slice)
	if len == 0 {
		return nil, errors.New("empty queue")
	}

	elementToReturn := st.slice[0]
	st.slice = st.slice[1:]

	return elementToReturn, nil
}

func (st *FIFOQueue) Get(index int) (interface{}, error) {
	st.RLock()
	defer st.RUnlock()

	if len(st.slice) <= index {
		return nil, errors.New(fmt.Sprintf("index out of bounds: %v", index))
	}

	return st.slice[index], nil
}

func (st *FIFOQueue) Remove(index int) error {
	st.Lock()
	defer st.Unlock()

	if len(st.slice) <= index {
		return errors.New(fmt.Sprintf("index out of bounds: %v", index))
	}

	st.slice = append(st.slice[:index], st.slice[index+1:]...)
	return nil
}

func (st *FIFOQueue) RemoveValue(value interface{}) error {
	st.Lock()
	defer st.Unlock()

	for i := 0; i < len(st.slice); i++ {
		if st.slice[i] == value {
			st.slice = append(st.slice[:i], st.slice[i+1:]...)
			i--
		}
	}
	return nil
}

func (st *FIFOQueue) GetLen() int {
	st.RLock()
	defer st.RUnlock()
	return len(st.slice)
}

func (st *FIFOQueue) DebugPrint(){
	fmt.Printf("slice len=%d - %v\n", len(st.slice), st.slice)
}

func (st *FIFOQueue) GetCap() int {
	st.RLock()
	defer st.RUnlock()
	return cap(st.slice)
}
