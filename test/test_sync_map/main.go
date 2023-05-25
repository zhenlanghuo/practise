package main

import (
	"fmt"
	"strings"
	"sync"
	//"container/heap"
)

type KV struct {
	key   string
	value string
}

func main() {
	m := sync.Map{}
	//m.Store(1, 1)
	//m.Store(2, 2)
	//m.Store(3, 3)
	//for i := 0; i < 4; i++ {
	//	m.Load(1)
	//}
	//m.Delete(3)
	//m.Store(4, 4)
	//m.Delete(2)

	//fmt.Println(m.Load(3))

	//fmt.Println(1)

	m.Store("test_string", &KV{key: "test_string", value: "test_string"})
	e2, _ := m.Load("test_string")
	fmt.Printf("%p, %p\n", e2, &e2)
	for i := 0; i < 10; i++ {
		m.Load("test_string")
	}
	fmt.Println(m.Load("test_string"))
	e, _ := m.Load("test_string")
	fmt.Printf("%p, %p\n", e, &e)
	(e.(*KV)).value = "aaa"
	//m.Store("test_string", e)
	m.Store("test_string", &KV{key: "test_string", value: "aaa"})
	fmt.Println(m.Load("test_string"))
	e1, _ := m.Load("test_string")
	fmt.Printf("%p, %p\n", e1, &e1)

	strings.Builder{}
}
