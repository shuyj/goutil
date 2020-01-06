package szcache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCache_Set(t *testing.T) {

	c := NewCacheBySize(10)

	for i:=0; i<20; i++ {
		s := fmt.Sprintf("%d", i)
		c.Set(s, s, 3*time.Second)
	}

	for i:=0; i<20; i++ {
		s := fmt.Sprintf("%d", i)
		v,t,e := c.GetWithTime(s)
		if e != nil {
			fmt.Printf("key=%s err=%s\n", s, e.Error())
			continue
		}
		fmt.Printf("key=%s val=%s time=%v\n", s, v, t)
	}

}

func BenchmarkCache_Set(b *testing.B) {
	c := NewCacheBySize(1000000)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			s := fmt.Sprintf("%d", rand.Intn(1000000))
			c.Set(s, s, 3*time.Second)
		}
	})
}

func BenchmarkCache_Get(b *testing.B) {
	c := NewCacheBySize(1000000)
	for i:=0; i<1000000; i++{
		s := fmt.Sprintf("%d", rand.Intn(1000000))
		c.Set(s, s, 3*time.Second)
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			s := fmt.Sprintf("%d", rand.Intn(1000000))
			c.GetWithTime(s)
		}
	})
}


func BenchmarkCache_SetGet(b *testing.B){
	gc := NewCacheBySize(10000)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := gc.Set("hello", "world", 3*time.Second)
			val, err := gc.Get("hello")
			if err != nil {
				b.Log("Get err=", err, val)
			}
		}
	})

}





