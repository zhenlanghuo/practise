package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"unsafe"
)

type minHeap []int

func (m minHeap) Len() int           { return len(m) }
func (m minHeap) Less(i, j int) bool { return m[i] < m[j] }
func (m minHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(int))
}

func (m *minHeap) Pop() interface{} {
	length := len(*m)
	result := (*m)[length-1]
	*m = (*m)[:length-1]
	return result
}

func main() {
	//var mh *minHeap

	mh := &minHeap{}

	heap.Init(mh)
	fmt.Println(mh)

	fmt.Printf("func s: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", mh, mh, cap(*mh), len(*mh), unsafe.Pointer(&mh))

	heap.Push(mh, rand.Intn(100))

	fmt.Printf("func s: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", mh, mh, cap(*mh), len(*mh), unsafe.Pointer(&mh))

	for i := 0; i < 100; i++ {
		heap.Push(mh, rand.Intn(100))
	}


	fmt.Println(mh)

	fmt.Printf("func s: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", mh, mh, cap(*mh), len(*mh), unsafe.Pointer(&mh))

}
