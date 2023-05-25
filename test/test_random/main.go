package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	n := 5
	m := make(map[int]int)
	for i:=0;i<n;i++{
		m[i] = 0
	}
	for i:=0;i<n*10;i++{
		m[rand.Intn(n)]++
	}
	fmt.Println(m)
}
