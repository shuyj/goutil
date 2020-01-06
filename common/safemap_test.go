package common

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkSMap_Put(b *testing.B) {
	m := NewSMap(10000)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			s := fmt.Sprintf("%d", rand.Intn(1000000))
			m.Put(s, s)
		}
	})
}

func BenchmarkSMap_Get(b *testing.B) {
	m := NewSMap(10000)
	for i:=0; i<100000; i++{
		s := fmt.Sprintf("%d", i)
		m.Put(s, s)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			s := fmt.Sprintf("%d", rand.Intn(100000))
			m.Get(s)
		}
	})
}

func BenchmarkSyncMap_Put(b *testing.B) {
	m := sync.Map{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			s := fmt.Sprintf("%d", rand.Intn(1000000))
			m.Store(s, s)
		}
	})
}

func BenchmarkSyncMap_Get(b *testing.B) {
	m := sync.Map{}
	for i:=0; i<100000; i++{
		s := fmt.Sprintf("%d", i)
		m.Store(s, s)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			s := fmt.Sprintf("%d", rand.Intn(100000))
			m.Load(s)
		}
	})
}
