package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//s := []int{1, 2, 3}

	s := make([]int, 0, 3)

	fmt.Printf("func s: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", s, s, cap(s), len(s), unsafe.Pointer(&s))

	a := append(s, 2)

	fmt.Printf("func s: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", s, s, cap(s), len(s), unsafe.Pointer(&s))
	fmt.Printf("func a: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", a, a, cap(a), len(a), unsafe.Pointer(&a))

	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 2)
	fmt.Printf("func s: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", s, s, cap(s), len(s), unsafe.Pointer(&s))
	fmt.Printf("func a: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", a, a, cap(a), len(a), unsafe.Pointer(&a))

	testSlice(s)


	//m := map[int]int{}

}

func testSlice(s []int) {
	fmt.Printf("testSlice s: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", s, s, cap(s), len(s), unsafe.Pointer(&s))
}
