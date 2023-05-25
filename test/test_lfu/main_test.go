package main

import "testing"

func BenchmarkLFULogIncr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LFULogIncr(125)
	}
}
