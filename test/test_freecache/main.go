package main

import "github.com/coocood/freecache"

func main() {
	cache := freecache.NewCache(10)
	cache.Set()

	//cache := bigcache.New()
	//cache.Set()

}
