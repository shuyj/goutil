package common

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkOnceParallel(b *testing.B){
	var once sync.Once

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			once.Do(func() {
				b.Logf("sleep done %p", &once)
			})
		}
	})

}

func BenchmarkNewFilterLock(b *testing.B){
	lck := NewNoBlockLock()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			go func() {
				if lck.Lock() {
					lck.Unlock()
				}
			}()
		}
	})
}

func BenchmarkSystemLock(b *testing.B){
	lck := sync.Mutex{}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			go func() {
				lck.Lock()
				lck.Unlock()
			}()
		}
	})
}

func BenchmarkNewNoBlockSem(b *testing.B){
	sem := NewNoBlockSem(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			go func() {
				if sem.Acquire() {
					sem.Release()
				}
			}()
		}
	})
}

func TestSliceIndex(t *testing.T){
	s := []int{0,1,8,2,3,8}
	fmt.Printf("slice origin length = %d \n", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Println(i, "value=", s[i])
		if s[i] == 8 {
			s = append(s[:i], s[i+1:]...)
			i-- // maintain the correct index
			fmt.Printf("slice change length = %d \n", len(s))
		}
	}

	fmt.Printf("slice = %v \n", s)
}

