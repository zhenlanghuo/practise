package main

import (
	"testing"
)

type A struct {
	a1 int
	b  *B
	a2 float64
	a3 float64
	a4 uint64
	s  string
}

type B struct {
	b1   int
	b2   float64
	data []byte
}

func Benchmark_GoAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := &Person{}
		p.a = 1
	}
}

func Benchmark_CgoAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//p, err := New[Person]()
		//if err != nil {
		//	fmt.Println("err", err)
		//	return
		//}
		//p.a = 1
		size := 858
		//var bytes []byte
		bytes := make([]byte, size)
		//sh := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
		//sh.Data = uintptr(jemalloc.Malloc(int(size)))
		//sh.Len = int(size)
		//sh.Cap = int(size)
		bytes[0] = 2
	}
}
