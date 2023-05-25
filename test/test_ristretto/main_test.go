package main

import (
	"github.com/dgraph-io/ristretto"
	"math/rand"
	"strconv"
	"testing"
)

func BenchmarkGet(b *testing.B) {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10000; i++ {
		// set a value with a cost of 1
		cache.Set(strconv.FormatInt(int64(i), 10), "value", 1)
	}
	cache.Wait()

	for i := 0; i < b.N; i++ {
		cache.Get(rand.Intn(10000))
	}
}
