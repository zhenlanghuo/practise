package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type tflag uint8
type nameOff int32 // offset to a name
type typeOff int32 // offset to an *rtype

type TestStruct struct {
	ptr uintptr
	f1  uint8
	f2  *uint8
	f3  uint32
	f4  *uint64
	f5  uint64
}

type rtype struct {
	size       uintptr
	ptrdata    uintptr // number of bytes in the type that can contain pointers
	hash       uint32  // hash of type; avoids computation in hash tables
	tflag      tflag   // extra type information flags
	align      uint8   // alignment of variable with this type
	fieldAlign uint8   // alignment of struct field with this type
	kind       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	equal     func(unsafe.Pointer, unsafe.Pointer) bool
	gcdata    *byte   // garbage collection data
	str       nameOff // string form
	ptrToThis typeOff // type for pointer to this type, may be zero
}

func main() {
	t := &TestStruct{}
	//t := make([]int, 0)
	//v := reflect.TypeOf(*t)
	//rtyp, ok := v.(*reflect.Rtype)
	//if !ok {
	//	fmt.Printf("error")
	//	return
	//}
	//
	//r := (*rtype)(unsafe.Pointer(rtyp))
	//
	//fmt.Printf("%#v\n", *r)
	//fmt.Printf("*gcdata=%d\n", *(r.gcdata))

	fmt.Println(reflect.ValueOf(t).Elem().NumField())

	//json.Marshal()
}
