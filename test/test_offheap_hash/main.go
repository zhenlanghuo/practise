package main

import (
	"github.com/allegro/bigcache/v3"
)

func main() {
	//table := offheap.NewStringHashTable(1024)
	//table.InsertStringKey()
	//cache := fastcache.New(100)
	//cache.Set()
	//cache.Get()
	cache := bigcache.New()
	cache.Set()
	cache.Get()
}
