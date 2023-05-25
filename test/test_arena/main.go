package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type Person struct {
	name  string
	bytes []byte
}

type sliceHeader struct {
	Data unsafe.Pointer
	Len  int64
	Cap  int64
}

type StringHeader struct {
	Data uintptr
	Len  int
}

// 假设此函数不会被内联（inline）。
func createInt() *int {
	return new(int)
}

func foo() {
	p0, y, z := createInt(), createInt(), createInt()
	var p1 = unsafe.Pointer(y) // 和y一样引用着同一个值
	var p2 = uintptr(unsafe.Pointer(z))

	// 此时，即使z指针值所引用的int值的地址仍旧存储
	// 在p2值中，但是此int值已经不再被使用了，所以垃圾
	// 回收器认为可以回收它所占据的内存块了。另一方面，
	// p0和p1各自所引用的int值仍旧将在下面被使用。

	// uintptr值可以参与算术运算。
	p2 += 2
	p2--
	p2--

	runtime.GC()

	*p0 = 1                         // okay
	*(*int)(p1) = 2                 // okay
	*(*int)(unsafe.Pointer(p2)) = 3 // 危险操作！
}

func main() {
	//a := arena.NewArena()
	//p := arena.New[Person](a)
	//p.name = "1213"
	//fmt.Println(p)
	////runtime.GC()
	//a.Free()
	//runtime.GC()
	//fmt.Println(p)

	//acPool := lac.NewAllocatorPool("test_string", nil, 10, 100*1024, 10, 10)
	//ac := acPool.Get()
	//defer ac.Release()

	for i := 0; i < 1000; i++ {
		foo()
	}
	fmt.Println(1111)

	//var p *Person
	//fmt.Println(unsafe.Sizeof(*p))
	//p = &Person{name: "test1", bytes: make([]byte, 1024*1024)}
	//p.bytes[1024*1024-1] = byte(rand.Intn(256))
	//puintptr1 := (uintptr)(unsafe.Pointer(p))
	//p = nil
	//
	//fmt.Println(1)
	//fmt.Println((*Person)(unsafe.Pointer(puintptr1)).name)
	//fmt.Println(len((*Person)(unsafe.Pointer(puintptr1)).bytes))
	//fmt.Println((*Person)(unsafe.Pointer(puintptr1)).bytes[1024*1024-1])
	//time.Sleep(time.Second * 1)
	//runtime.GC()
	//debug.FreeOSMemory()
	//time.Sleep(time.Second * 1)
	//fmt.Println(2)
	//fmt.Println((*Person)(unsafe.Pointer(puintptr1)).name)
	//fmt.Println(len((*Person)(unsafe.Pointer(puintptr1)).bytes))
	//fmt.Println((*Person)(unsafe.Pointer(puintptr1)).bytes[1024*1024-1])
	//fmt.Println((*Person)(unsafe.Pointer(puintptr1)))

	//array := make([]byte, 64)
	//fmt.Println(array)
	//arrayPointer := unsafe.Pointer(&array)
	//arrayHeader := (*sliceHeader)(arrayPointer)
	//
	//arrayDataUintptr := uintptr(arrayHeader.Data)
	//fmt.Println(*arrayHeader)
	//fmt.Println(arrayDataUintptr)
	//fmt.Println((*[64]byte)(unsafe.Pointer(arrayDataUintptr)))
	//
	//pPointer := unsafe.Add(arrayHeader.Data, unsafe.Sizeof(*p))
	//pPointerUintptr := uintptr(pPointer)
	//fmt.Println(pPointer, pPointerUintptr)
	//p = (*Person)(pPointer)
	////p = lac.New[Person](ac)
	////p = lac.New[Person](ac)
	//p.name = "1213"
	//s1213Uintptr := (*StringHeader)(unsafe.Pointer(&p.name)).Data
	//
	//fmt.Println((*(*[4]byte)(unsafe.Pointer(s1213Uintptr))))
	//fmt.Println(array)
	//fmt.Println((*[64]byte)(unsafe.Pointer(arrayDataUintptr)))
	//fmt.Println(p)
	//p.name = "2342352"
	//fmt.Println(array)
	//fmt.Println((*[64]byte)(unsafe.Pointer(arrayDataUintptr)))
	//fmt.Println(p)
	////ac.Release()
	//array = nil
	//p = nil
	//runtime.GC()
	//runtime.GC()
	//runtime.GC()
	//runtime.GC()
	//runtime.GC()
	//runtime.GC()
	//fmt.Println(array)
	//fmt.Println(*arrayHeader)
	//fmt.Println(arrayDataUintptr)
	//fmt.Println((*[64]byte)(unsafe.Pointer(arrayDataUintptr)))
	////fmt.Println(array[0])
	////fmt.Println(array[16])
	////runtime.GC()
	//fmt.Println(*(*[4]byte)(unsafe.Pointer(s1213Uintptr)))
	//fmt.Println(p)
}
