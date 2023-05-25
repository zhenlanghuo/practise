//package main
//
//import (
//	"fmt"
//	"reflect"
//	"runtime"
//	"time"
//	"unsafe"
//)
//
//func main() {
//	s := "1122331"
//	suintptr1 := (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
//	suintptr1Len := (*reflect.StringHeader)(unsafe.Pointer(&s)).Len
//	//fmt.Println(suintptr1)
//	//fmt.Println(uintptrToString(suintptr1, suintptr1Len))
//	s = "223344"
//	suintptr2 := (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
//	suintptr2Len := (*reflect.StringHeader)(unsafe.Pointer(&s)).Len
//	fmt.Println(suintptr2)
//	fmt.Println(uintptrToString(suintptr2, suintptr2Len))
//
//	//fmt.Println(suintptr1)
//	//fmt.Println(uintptrToString(suintptr1, suintptr1Len))
//	var array []byte
//	for i := 0; i < 100; i++ {
//		array = append(array, make([]byte, 1024*1024*10)...)
//		fmt.Println(array[0])
//	}
//	time.Sleep(time.Second * 10)
//	runtime.GC()
//	//runtime.GC()
//	//runtime.GC()
//	//runtime.GC()
//	//runtime.GC()
//	//runtime.GC()
//	//runtime.GC()
//	time.Sleep(time.Second * 1)
//	fmt.Println(suintptr1)
//	fmt.Println(uintptrToString(suintptr1, suintptr1Len))
//	fmt.Println(suintptr2)
//	fmt.Println(uintptrToString(suintptr2, suintptr2Len))
//}
//
//func uintptrToString(u uintptr, len int) string {
//	//fmt.Println(u)
//	bh := reflect.SliceHeader{
//		Data: u,
//		Len:  len,
//		Cap:  len,
//	}
//	return string(*(*[]byte)(unsafe.Pointer(&bh)))
//}

package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

type Person struct {
	name string
}

func main() {
	//fmt.Printf("%s\n", test())
	ptr := test1()
	//runtime.GC()
	//time.Sleep(time.Millisecond * 50)
	fmt.Printf("%v\n", (*Person)(unsafe.Pointer(ptr)))
	//fmt.Printf("12 %v\n", (*Person)(unsafe.Pointer(ptr)).name)
}

func test1() (ptr uintptr) {
	//defer runtime.GC()
	p := &Person{name: "test1"}
	ptr = uintptr(unsafe.Pointer(p))
	runtime.GC()
	//return (*Person)((unsafe.Pointer((uintptr)(unsafe.Pointer(p)))))
	return
}

//func test() []byte {
//	//defer runtime.GC()
//	x := make([]byte, 5)
//	x[0] = 'h'
//	x[1] = 'e'
//	x[2] = 'l'
//	x[3] = 'l'
//	x[4] = 'o'
//	return StringToSliceByte1(string(x))
//}
//
//func StringToSliceByte1(s string) []byte {
//	l := len(s)
//	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
//		Data: (*(*reflect.StringHeader)(unsafe.Pointer(&s))).Data,
//		Len:  l,
//		Cap:  l,
//	}))
//}
