package main

import (
	"fmt"
	"runtime"
	"sync"
)

type Node struct {
	next  *Node
	bytes []byte
}

func allocHeap() (*Node, *Node) {
	a, b := &Node{}, &Node{}
	a.bytes = make([]byte, 1024*1024*100)
	b.bytes = make([]byte, 1024*1024*100)
	a.next = b
	b.next = a
	// 注册 finalizer，输出信息
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	runtime.SetFinalizer(a, func(node *Node) {
		fmt.Println("Object a has been garbage collected.")
	})
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	runtime.SetFinalizer(b, func(node *Node) {
		fmt.Println("Object b has been garbage collected.")
	})
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	return a, b
}

func main() {
	//fmt.Println("main NumGoroutine", runtime.NumGoroutine())
	//go allocHeap()
	//fmt.Println("main NumGoroutine", runtime.NumGoroutine())
	//// 手动触发垃圾回收
	//for {
	//	runtime.GC()
	//	fmt.Println("GC done")
	//	fmt.Println("main NumGoroutine", runtime.NumGoroutine())
	//	time.Sleep(time.Second)
	//}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		wg.Done()
	}()

	wg.Wait()
	runtime.NumGoroutine()
	//fmt.Println("main NumGoroutine", runtime.NumGoroutine())
}
