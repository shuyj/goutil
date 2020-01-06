package common

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewFIFOQueue(t *testing.T) {

	q := NewFIFOQueue()

	for i:=0; i<20; i++{
		q.Enqueue(fmt.Sprintf("%d", i))
	}

	q.DebugPrint()
	q.RemoveValue("10")
	q.RemoveValue("15")

	for{
		v,e := q.Dequeue()
		if e != nil {
			break
		}
		fmt.Printf("dequueue value = %v\n", v)
	}

}

func BenchmarkFIFOQueue_Enqueue(b *testing.B) {
	q := NewFIFOQueue()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			q.Enqueue(fmt.Sprintf("%d", rand.Intn(10000000)))
		}
	})
}

func BenchmarkFIFOQueue_Dequeue(b *testing.B) {
	q := NewFIFOQueue()
	for i:=0; i<10000000; i++{
		q.Enqueue(fmt.Sprintf("%d", i))
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			q.Dequeue()
		}
	})
}

