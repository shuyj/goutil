package common

import (
	"testing"
	"time"
)

func TestNewOpenLock(t *testing.T) {

	lck := NewOpenLock()

	t.Parallel()
	AcquireForOneNoBlock(lck, func() bool {
		return true
	}, func() {
		t.Logf("exec %v", lck.Version)
	})

}

func BenchmarkNewOpenLock(b *testing.B) {
	prevTime := time.Now().Add(-time.Second*2)
	lck := NewOpenLock()
	b.Logf("Benchmark lck = %p", lck)
	b.RunParallel(func(pb *testing.PB) {
		b.Logf("runParallel %v lck = %p", pb, lck)
		for pb.Next() {
			AcquireForOneNoBlock(lck, func() bool {
				now := time.Now()
				if prevTime.Add(1*time.Second).Before(now) {
					prevTime = now
					return true
				}
				return false
			}, func() {
				b.Logf("%v  prev = %v exec verison = %p", time.Now(), prevTime, lck)
			})
		}
	})
}

func BenchmarkNewNoBlockLock(b *testing.B) {
	prevTime := time.Now().Add(-time.Second*2)
	lck := NewNoBlockLock()
	b.Logf("Benchmark lck = %p", lck)
	b.RunParallel(func(pb *testing.PB) {
		//b.Logf("runParallel %v lck = %p", pb, lck)
		for pb.Next() {
			AcquireForOneNoBlockLock(lck, func() bool {
				now := time.Now()
				if prevTime.Add(1*time.Second).Before(now) {
					prevTime = now
					return true
				}
				return false
			}, func() {
				b.Logf("%v  prev = %v exec verison = %p", time.Now(), prevTime, lck)
			}, 1*time.Second)
		}
	})
}

