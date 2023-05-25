package main

import (
	"github.com/crazybie/linear_ac/lac"
	"unsafe"
)

var ac *lac.Allocator

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

type PbItem struct {
	Id     *int
	Price  *int
	Class  *int
	Name   *string
	Active *bool
}

type PbData struct {
	Age   *int
	Items []*PbItem
	InUse *PbItem
}

var size = int(1e7)

func main() {
	acPool := lac.NewAllocatorPool("test_string", nil, 100, 100*1024*1024, 100, 100)
	ac = acPool.Get()
	defer ac.Release()

	//runtime.GOMAXPROCS(4)

	//s := testWithRawSlice()
	//defer runtime.KeepAlive(s)

	//m := testWithRaw()
	//defer runtime.KeepAlive(m)

	//m, s := testWithRawMaoAndSlice()
	//defer runtime.KeepAlive(m)
	//defer runtime.KeepAlive(s)

	//m, s := testWithLacMapAndSlice()
	//defer runtime.KeepAlive(m)
	//defer runtime.KeepAlive(s)

	//m := testWithLac()
	//defer runtime.KeepAlive(m)

	//for i := 0; i < 10; i++ {
	//	start := time.Now()
	//	runtime.GC()
	//	fmt.Printf("GC took %s\n", time.Since(start))
	//}

	//v := reflect.TypeOf(*m[1])
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
}

func testWithLac() map[int]*PbData {
	m := make(map[int]*PbData, size)
	for k := 0; k < size; k++ {
		d := lac.New[PbData](ac)
		d.Age = ac.Int(11)

		n := 3
		for i := 0; i < n; i++ {
			item := lac.New[PbItem](ac)
			item.Id = ac.Int(i + 1)
			item.Active = ac.Bool(true)
			item.Price = ac.Int(100 + i)
			item.Class = ac.Int(3 + i)
			item.Name = ac.String("name")

			d.Items = lac.Append(ac, d.Items, item)
		}
		m[k] = d
	}

	return m
}

func testWithLacMapAndSlice() (map[int]int, []*PbData) {
	m := make(map[int]int, size)
	s := make([]*PbData, size)
	for k := 0; k < size; k++ {
		d := lac.New[PbData](ac)
		d.Age = ac.Int(11)

		n := 3
		for i := 0; i < n; i++ {
			item := lac.New[PbItem](ac)
			item.Id = ac.Int(i + 1)
			item.Active = ac.Bool(true)
			item.Price = ac.Int(100 + i)
			item.Class = ac.Int(3 + i)
			item.Name = ac.String("name")

			d.Items = lac.Append(ac, d.Items, item)
		}
		m[k] = k
		s[k] = d
	}

	return m, s
}

func testWithRaw() map[int]*PbData {
	m := make(map[int]*PbData, size)
	for k := 0; k < size; k++ {
		age := 1

		d := &PbData{
			Age: &age,
		}

		n := 3
		for i := 0; i < n; i++ {
			id := i + 1
			active := true
			price := 100 + i
			class := 3 + i
			name := "name"

			item := &PbItem{
				Id:     &id,
				Active: &active,
				Price:  &price,
				Class:  &class,
				Name:   &name,
			}

			d.Items = append(d.Items, item)
		}
		m[k] = d
	}

	return m
}

func testWithRawInt() map[int]*int {
	m := make(map[int]*int, size)
	for k := 0; k < size; k++ {
		a := k
		m[k] = &a
	}
	return m
}

func testWithRawSlice() []*PbData {
	s := make([]*PbData, size)
	for k := 0; k < size; k++ {
		age := 1

		d := &PbData{
			Age: &age,
		}

		n := 3
		for i := 0; i < n; i++ {
			id := i + 1
			active := true
			price := 100 + i
			class := 3 + i
			name := "name"

			item := &PbItem{
				Id:     &id,
				Active: &active,
				Price:  &price,
				Class:  &class,
				Name:   &name,
			}

			d.Items = append(d.Items, item)
		}
		s[k] = d
	}

	return s
}

func testWithRawMaoAndSlice() (map[int]int, []*PbData) {
	m := make(map[int]int, size)
	s := make([]*PbData, size)
	for k := 0; k < size; k++ {
		age := 1

		d := &PbData{
			Age: &age,
		}

		n := 3
		for i := 0; i < n; i++ {
			id := i + 1
			active := true
			price := 100 + i
			class := 3 + i
			name := "name"

			item := &PbItem{
				Id:     &id,
				Active: &active,
				Price:  &price,
				Class:  &class,
				Name:   &name,
			}

			d.Items = append(d.Items, item)
		}
		s[k] = d
		m[k] = k
	}

	return m, s
}
