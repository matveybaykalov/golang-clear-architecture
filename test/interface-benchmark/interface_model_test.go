package interfacebenchmark_test

import (
	"fmt"
	"testing"
	"time"
)

type A struct {
	data int
}

func (a *A) inc() {
	a.data += 1
}

type B struct {
	data int
}

func (b *B) inc() {
	b.data += 1
}

type C interface {
	inc()
}

var l = 40000000

func BenchmarkStatic(b *testing.B) {

	arr1 := make([]A, l)
	arr2 := make([]B, l)

	b.ResetTimer()
	b.StartTimer()
	start := time.Now()
	for i := 0; i < l; i++ {
		arr1[i].inc()
	}

	for i := 0; i < l; i++ {
		arr2[i].inc()
	}
	stop := time.Now()
	b.StopTimer()
	fmt.Println(stop.Sub(start))
}

func BenchmarkDynamic(b *testing.B) {

	arr11 := make([]C, l)
	arr12 := make([]C, l)

	arr1 := make([]A, l)
	arr2 := make([]B, l)

	for i := 0; i < l; i++ {
		arr11[i] = &arr1[i]
		arr12[i] = &arr2[i]
	}

	b.ResetTimer()
	b.StartTimer()
	start := time.Now()
	for i := 0; i < l; i++ {
		arr11[i].inc()
	}

	for i := 0; i < l; i++ {
		arr12[i].inc()
	}
	stop := time.Now()
	b.StopTimer()
	fmt.Println(stop.Sub(start))
}
