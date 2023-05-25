package main

import (
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"reflect"
	"strconv"
	"unsafe"
)

type TestStruct struct {
	IntPointField     *int
	Float32PointField *float32
	Person            *Person
	IntField          int
	IntSlice          []int
	Persons           []*Person
}

func (t *TestStruct) String() string {

	if t == nil {
		return "{}"
	}

	IntPointFieldStr := "nil"
	if t.IntPointField != nil {
		IntPointFieldStr = strconv.FormatInt(int64(*t.IntPointField), 10)
	}

	Float32PointFieldStr := "nil"
	if t.Float32PointField != nil {
		Float32PointFieldStr = fmt.Sprintf("%f", *t.Float32PointField)
	}

	return fmt.Sprintf("IntPointField: %v, Float32PointField: %v, Person: %v, IntField: %v, IntSlice: %v, Persons: %v",
		IntPointFieldStr, Float32PointFieldStr, t.Person, t.IntField, t.IntSlice, t.Persons)
	//return ""
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	if p == nil {
		return "{}"
	}

	return fmt.Sprintf("{Name: %v, Age: %v}", p.Name, p.Age)
}

var m = make(map[interface{}]uint64)

func ToBytesSize(i interface{}) uint64 {
	if i == nil {
		return 0
	}

	size := uint64(0)

	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)

	switch t.Kind() {
	case reflect.Pointer:
		if size, ok := m[i]; ok {
			return size
		}
		v = v.Elem()
		t = t.Elem()
		size += uint64(t.Size())
		//fmt.Println(t.Name(), "t.Kind() == reflect.Pointer", uint64(t.Size()), size)
	case reflect.Struct:
		size += uint64(t.Size())
		//fmt.Println(t.Name(), "t.Kind() == reflect.Struct", uint64(t.Size()), size)
	case reflect.Slice:
		switch t.Elem().Kind() {
		case reflect.Pointer, reflect.Slice:
			size += uint64(v.Len()) * uint64(t.Elem().Size())
			//fmt.Println(t.Name(), "t.Kind() == reflect.Slice, t.Elem().Kind() == reflect.Pointer || reflect.Slice", uint64(v.Len())*uint64(t.Elem().Size()), size)
			for k := 0; k < v.Len(); k++ {
				size += ToBytesSize(v.Index(k).Interface())
				//fmt.Println(t.Name(), "t.Kind() == reflect.Slice, t.Elem().Kind() == reflect.Pointer || reflect.Slice, k:", k, ToBytesSize(v.Index(k).Interface()), size)
			}
		case reflect.Struct:
			for k := 0; k < v.Len(); k++ {
				size += ToBytesSize(v.Index(k).Interface())
				//fmt.Println(t.Name(), "t.Kind() == reflect.Slice, t.Elem().Kind() == reflect.Struct, k:", k, ToBytesSize(v.Index(k).Interface()), size)
			}
		case reflect.Map:
			panic("not support map in slice")
		default:
			size += uint64(v.Len()) * uint64(t.Elem().Size())
			//fmt.Println(t.Name(), "t.Kind() == default", uint64(v.Len())*uint64(t.Elem().Size()), size)
		}
		return size
	}

	for fieldIndex := 0; fieldIndex < v.NumField(); fieldIndex++ {
		fieldValue := v.Field(fieldIndex)
		fieldType := fieldValue.Type()
		if !t.Field(fieldIndex).IsExported() {
			continue
		}
		//fmt.Println(t.Field(fieldIndex).Name, fieldValue.String(), fieldValue.Type(), fieldValue, fieldType.Kind())
		switch fieldType.Kind() {
		case reflect.Pointer:
			if fieldValue.IsNil() {
				continue
			}
			switch fieldType.Elem().Kind() {
			case reflect.Struct:
				size += ToBytesSize(fieldValue.Interface())
			//fmt.Println("fieldType.Elem().Kind() is reflect.Struct", ToBytesSize(fieldValue.Interface()), size)
			case reflect.String:
				size += uint64(fieldType.Elem().Size())
				size += uint64(fieldValue.Elem().Len())
			default:
				size += uint64(unsafe.Sizeof(fieldType.Elem().Size()))
				//fmt.Println("fieldType.Elem().Kind() is not reflect.Struct", uint64(unsafe.Sizeof(fieldType.Elem().Size())), size)
			}
		case reflect.String:
			if fieldValue.IsZero() {
				continue
			}
			size += uint64(fieldValue.Len())
			//fmt.Println("fieldType.Kind() is string", uint64(fieldValue.Len()), size)
		case reflect.Slice:
			size += ToBytesSize(fieldValue.Interface())
			//fmt.Println("fieldType.Kind() is slice", ToBytesSize(fieldValue.Interface()), size)
		}
	}
	m[i] = size
	return size
}

func (t *TestStruct) ToBytes(uptr *uintptr) ([]byte, []uint64) {
	var bytes []byte
	var pointIndex []uint64
	pointIndex = make([]uint64, 0)

	if uptr == nil {
		bytes = make([]byte, ToBytesSize(t))
		uptr = new(uintptr)
		*uptr = (*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data
		//pointIndex = append(pointIndex, uint64(*uptr))
	}
	newT := (*TestStruct)(unsafe.Pointer(*uptr))
	*uptr += unsafe.Sizeof(*t)

	if t.IntPointField != nil {
		newT.IntPointField = (*int)(unsafe.Pointer(*uptr))
		*newT.IntPointField = *t.IntPointField
		*uptr += unsafe.Sizeof(*t.IntPointField)
		pointIndex = append(pointIndex, uint64(uintptr(unsafe.Pointer(&newT.IntPointField))))
	}

	if t.Float32PointField != nil {
		newT.Float32PointField = (*float32)(unsafe.Pointer(*uptr))
		*newT.Float32PointField = *t.Float32PointField
		*uptr += unsafe.Sizeof(*t.Float32PointField)
		pointIndex = append(pointIndex, uint64(uintptr(unsafe.Pointer(&newT.Float32PointField))))
	}

	if t.Person != nil {
		newT.Person = (*Person)(unsafe.Pointer(*uptr))
		_, pointIndex_ := t.Person.ToBytes(uptr)
		pointIndex = append(pointIndex, uint64(uintptr(unsafe.Pointer(&newT.Person))))
		pointIndex = append(pointIndex, pointIndex_...)
	}

	newT.IntField = t.IntField

	if len(t.IntSlice) != 0 {
		slice := (*reflect.SliceHeader)(unsafe.Pointer(&newT.IntSlice))
		slice.Cap = len(t.IntSlice)
		slice.Len = len(t.IntSlice)
		slice.Data = *uptr
		pointIndex = append(pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
		for index, v := range t.IntSlice {
			//fmt.Println("!@#", index, v)
			newT.IntSlice[index] = v
		}
		*uptr += uintptr(int(unsafe.Sizeof(int(0))) * len(t.IntSlice))
	}

	if len(t.Persons) != 0 {
		slice := (*reflect.SliceHeader)(unsafe.Pointer(&newT.Persons))
		slice.Cap = len(t.Persons)
		slice.Len = len(t.Persons)
		slice.Data = *uptr
		pointIndex = append(pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
		*uptr += uintptr(int(unsafe.Sizeof(int(0))) * len(t.Persons))
		for index, v := range t.Persons {
			newT.Persons[index] = (*Person)(unsafe.Pointer(*uptr))
			pointIndex = append(pointIndex, uint64(uintptr(unsafe.Pointer(&newT.Persons[index]))))
			_, pointIndex_ := v.ToBytes(uptr)
			pointIndex = append(pointIndex, pointIndex_...)
		}
	}

	return bytes, pointIndex
}

func (p *Person) ToBytes(uptr *uintptr) ([]byte, []uint64) {
	var bytes []byte
	var pointIndex []uint64
	pointIndex = make([]uint64, 0)

	if uptr == nil {
		bytes = make([]byte, ToBytesSize(p))
		*uptr = (*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data
	}

	newP := (*Person)(unsafe.Pointer(*uptr))
	*uptr += unsafe.Sizeof(*newP)
	if len(p.Name) != 0 {
		slice := (*reflect.StringHeader)(unsafe.Pointer(&newP.Name))
		slice.Len = len(p.Name)
		slice.Data = *uptr
		pointIndex = append(pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
		strbytes := *(*[]byte)(unsafe.Pointer(&newP.Name))
		for index, v := range p.Name {
			strbytes[index] = byte(v)
		}
		*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(p.Name))
	}

	newP.Age = p.Age

	return bytes, pointIndex
}

const intWidth int = int(unsafe.Sizeof(0))

var byteOrder binary.ByteOrder

// ByteOrder returns the byte order for the CPU's native endianness.
func ByteOrder() binary.ByteOrder { return byteOrder }
func init() {
	i := int(0x1)
	if v := (*[intWidth]byte)(unsafe.Pointer(&i)); v[0] == 0 {
		byteOrder = binary.BigEndian
	} else {
		byteOrder = binary.LittleEndian
	}
}

func main() {

	//fmt.Println(byteOrder)
	//
	//var ts *TestStruct
	//ts = &TestStruct{}

	//fmt.Printf("%p\n", &ts)
	//fmt.Printf("%v\n", &ts)
	//fmt.Printf("%p\n", ts)
	//fmt.Printf("%v\n", ts)
	//fmt.Printf("%p\n", &ts.IntPointField)
	//fmt.Printf("%v\n", ts.IntPointField)
	//fmt.Printf("%p\n", &ts.Float32PointField)
	//fmt.Printf("%v\n", ts.Float32PointField)

	//bytes := make([]byte, unsafe.Sizeof(*ts))
	//ts = (*TestStruct)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data))
	//num := 100
	//ts.IntPointField = &num
	//
	////fmt.Printf("%v\n", ts)
	////fmt.Printf("%p\n", ts)
	////fmt.Printf("%v\n", (*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data)
	////fmt.Printf("%v\n", (uint64)((*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data))
	////fmt.Printf("%v\n", strconv.FormatUint((uint64)((*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data), 16))
	//
	//fmt.Printf("%p\n", ts.IntPointField)
	//fmt.Printf("%v\n", ts.IntPointField)
	//
	//hex := fmt.Sprintf("%v", ts.IntPointField)[2:]
	//n, _ := strconv.ParseUint(hex, 16, 64)
	//fmt.Printf("%v\n", n)
	//bytes_ := make([]byte, 8)
	//byteOrder.PutUint64(bytes_, n)
	//fmt.Printf("%v\n", bytes_)
	//fmt.Printf("%v\n", bytes)
	//fmt.Printf("%v\n", byteOrder.Uint64(bytes))
	//

	//num_ := *(*int)(unsafe.Pointer(uintptr(n)))
	//fmt.Printf("%v", num_)

	//var ts *TestStruct
	//fmt.Printf("%v\n", ts.ToBytesSize())
	num := 100
	ts := &TestStruct{IntPointField: &num, Float32PointField: proto.Float32(1.98), Person: &Person{Name: "123", Age: 1}, IntSlice: []int{1, 2, 3}, Persons: []*Person{&Person{Name: "12345", Age: 2}}}
	//ts := &TestStruct{Person: &Person{Name: "123", Age: 1}}
	fmt.Printf("%v\n", ToBytesSize(ts))
	fmt.Println(ts.String())

	bytes, pointIndex := ts.ToBytes(nil)
	fmt.Println(bytes)
	fmt.Println(pointIndex)

	newts := (*TestStruct)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data))
	fmt.Println(ts.String())
	fmt.Println(newts.String())

	clone := make([]byte, len(bytes))
	copy(clone, bytes)
	startAddr := uint64((*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data)
	cloneStartAddr := uint64((*reflect.SliceHeader)(unsafe.Pointer(&clone)).Data)
	for _, addr := range pointIndex {
		delt := byteOrder.Uint64(bytes[int(addr-startAddr):int(addr-startAddr)+8]) - startAddr
		byteOrder.PutUint64(clone[int(addr-startAddr):int(addr-startAddr)+8], cloneStartAddr+delt)
	}

	newts2 := (*TestStruct)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&clone)).Data))

	fmt.Println(bytes)
	fmt.Println(pointIndex)
	fmt.Println(clone)

	fmt.Println(ts.String())
	fmt.Println(newts.String())
	fmt.Println(newts2.String())

	for i := 0; i < len(bytes); i++ {
		bytes[i] = 0
	}

	fmt.Println(ts.String())
	fmt.Println(newts.String())
	fmt.Println(newts2.String())

	//fmt.Printf("%v\n", ToBytesSize(ts))
	//fmt.Printf("%v\n", unsafe.Sizeof(ts))
	//fmt.Printf("%v\n", unsafe.Sizeof(*ts))
	//ts.IntPointField = &num
	//fmt.Printf("%v\n", ts.ToBytesSize())

}
