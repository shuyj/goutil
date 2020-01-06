package acacheimpl

import (
	"testing"
	"time"
)

func TestACache(t *testing.T){
	gc := NewCacheByACache(10*time.Minute)
	err := gc.Set("hello", "world")
	val, err := gc.Get("hello")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(val)
}

func BenchmarkACache(b *testing.B){
	gc := NewCacheByACache(10*time.Minute)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := gc.Set("hello", "world")
			val, err := gc.Get("hello")
			if err != nil {
				b.Log("Get err=", err, val)
			}
		}
	})
}