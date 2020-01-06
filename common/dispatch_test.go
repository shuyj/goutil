package common

import (
	"testing"
)

func func1()func(){
	v := 0

	return func() {
		v++
		//fmt.Println(v)
	}
}

func TestAsync(t *testing.T) {

	ftest := func1()

	Async(ftest)


	//time.Sleep(2*time.Second)
}

func BenchmarkAsync(b *testing.B) {

	ftest := func1()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next(){
			Async(ftest)
		}
	})

}

