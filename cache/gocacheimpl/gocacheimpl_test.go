package gocacheimpl

import (
	"testing"
	"time"
)

func TestGoCache(t *testing.T){
	gc := NewCacheByGoCache(10*time.Minute)
	err := gc.Set("hello", "world")
	t.Log("Set err=", err)
	val, err := gc.Get("hello")
	if err != nil {
		t.Error("Get err=", err)
	}
	t.Log("Get result=",val)
}




func BenchmarkGoCache(b *testing.B){
	gc := NewCacheByGoCache(10*time.Minute)
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
