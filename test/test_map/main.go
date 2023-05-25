package main

import (
	"fmt"
	"unsafe"
)

type Test struct {
	a []int
	b string
}

func main() {

	a := [2]int{1}
	b := [2]int{2, 2}
	if a == b {

	}

	//t1 := Test{a: [1]int{1}, b: "1"}
	//t2 := Test{a: [1]int{1}, b: "1"}
	//
	//if t1 == t2 {
	//	fmt.Println("t1 == t2")
	//}
	m := make(map[int]int)
	fmt.Printf("main %v, addr=%p, addr(&m)=%p, unsafe.Pointer: %v\n", m, m, &m, unsafe.Pointer(&m))
	testMap(m)
	fmt.Printf("main %v, addr=%p, addr(&m)=%p, unsafe.Pointer: %v\n", m, m, &m, unsafe.Pointer(&m))


	t := &Test{}
	fmt.Printf("main %v, addr=%p, addr(*)=%p, addr(&t)=%p, unsafe.Pointer: %v\n", t, t, &*t, &t, unsafe.Pointer(&t))
	testPointer(t)
}

func testMap(m map[int]int) {
	fmt.Printf("testMap %v, addr=%p, addr(&m)=%p, unsafe.Pointer: %v\n", m, m, &m, unsafe.Pointer(&m))
	m[1] = 1
}

func testPointer(t *Test) {
	fmt.Printf("main %v, addr=%p, addr(*)=%p, addr(&t)=%p, unsafe.Pointer: %v\n", t, t, &*t, &t, unsafe.Pointer(&t))
}
