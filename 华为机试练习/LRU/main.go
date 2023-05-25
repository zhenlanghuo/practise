package main

import (
	"container/list"
	"sync"
)

type Solution struct {
	l        *list.List
	m        map[int]*list.Element
	capacity int
	mu       sync.Mutex
}

type KV struct {
	key   int
	value int
}

func Constructor(capacity int) Solution {
	return Solution{
		l:        list.New(),
		m:        make(map[int]*list.Element),
		capacity: capacity,
		mu:       sync.Mutex{},
	}
}

func (this *Solution) get(key int) int {

	this.mu.Lock()
	defer this.mu.Unlock()

	elem, ok := this.m[key]
	if ok {
		this.l.MoveToFront(elem)
		return elem.Value.(*KV).value
	}

	return 0
}

func (this *Solution) set(key int, value int) {
	// write code here

	this.mu.Lock()
	defer this.mu.Unlock()

	elem, ok := this.m[key]
	if ok {
		this.l.MoveToFront(elem)
		elem.Value.(*KV).value = value
		return
	}

	if this.l.Len() == this.capacity {
		back := this.l.Back()
		delete(this.m, back.Value.(*KV).key)
		this.l.Remove(back)
	}

	this.m[key] = this.l.PushFront(&KV{key: key, value: value})
}
