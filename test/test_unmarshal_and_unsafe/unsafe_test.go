package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"practise/test/test_unmarshal_and_unsafe/pb"
	"reflect"
	"testing"
	"unsafe"
)

func PbTestStructToBytesSize(t *pb.TestStruct) (size uint64, pointSize uint64) {
	if t == nil {
		return
	}

	size += uint64(unsafe.Sizeof(*t))
	pointSize += 1

	//if t.IntPointField != nil {
	//	size += uint64(unsafe.Sizeof(*t.IntPointField))
	//	pointSize += 1
	//}
	//if t.FloatPointField != nil {
	//	size += uint64(unsafe.Sizeof(*t.FloatPointField))
	//	pointSize += 1
	//}

	if t.Person != nil {
		size_, pointSize_ := PbPersonToBytesSize(t.Person)
		size += size_
		pointSize += pointSize_
	}

	//if t.IntField != nil {
	//	size += uint64(unsafe.Sizeof(*t.IntField))
	//	pointSize += 1
	//}

	if len(t.IntSlice) != 0 {
		size += uint64(len(t.IntSlice)) * uint64(unsafe.Sizeof(t.IntSlice[0]))
		pointSize += 1
	}

	for i := 0; i < len(t.Persons); i++ {
		if t.Persons[i] != nil {
			size_, pointSize_ := PbPersonToBytesSize(t.Persons[i])
			size += size_ + 8
			pointSize += pointSize_
		}
	}
	pointSize += 1

	//if t.I1 != nil {
	//	size += uint64(unsafe.Sizeof(*t.I1))
	//	pointSize += 1
	//}
	//if t.I2 != nil {
	//	size += uint64(unsafe.Sizeof(*t.I2))
	//	pointSize += 1
	//}
	//if t.I3 != nil {
	//	size += uint64(unsafe.Sizeof(*t.I3))
	//	pointSize += 1
	//}
	//if t.I4 != nil {
	//	size += uint64(unsafe.Sizeof(*t.I4))
	//	pointSize += 1
	//}
	//if t.I5 != nil {
	//	size += uint64(unsafe.Sizeof(*t.I5))
	//	pointSize += 1
	//}

	if len(t.S1) != 0 {
		size += uint64(len(t.S1)) * uint64(unsafe.Sizeof(byte(0)))
		pointSize += 1
	}
	if len(t.S2) != 0 {
		size += uint64(len(t.S2)) * uint64(unsafe.Sizeof(byte(0)))
		pointSize += 1
	}
	if len(t.S3) != 0 {
		size += uint64(len(t.S3)) * uint64(unsafe.Sizeof(byte(0)))
		pointSize += 1
	}
	if len(t.S4) != 0 {
		size += uint64(len(t.S4)) * uint64(unsafe.Sizeof(byte(0)))
		pointSize += 1
	}
	if len(t.S5) != 0 {
		size += uint64(len(t.S5)) * uint64(unsafe.Sizeof(byte(0)))
		pointSize += 1
	}

	//if t.S1 != nil {
	//	size += uint64(unsafe.Sizeof(*t.S1))
	//	pointSize += 1
	//	size += uint64(len(*t.S1)) * uint64(unsafe.Sizeof(byte(0)))
	//	pointSize += 1
	//}
	//if t.S2 != nil {
	//	size += uint64(unsafe.Sizeof(*t.S2))
	//	pointSize += 1
	//	size += uint64(len(*t.S2)) * uint64(unsafe.Sizeof(byte(0)))
	//	pointSize += 1
	//}
	//if t.S3 != nil {
	//	size += uint64(unsafe.Sizeof(*t.S3))
	//	pointSize += 1
	//	size += uint64(len(*t.S3)) * uint64(unsafe.Sizeof(byte(0)))
	//	pointSize += 1
	//}
	//if t.S4 != nil {
	//	size += uint64(unsafe.Sizeof(*t.S4))
	//	pointSize += 1
	//	size += uint64(len(*t.S4)) * uint64(unsafe.Sizeof(byte(0)))
	//	pointSize += 1
	//}
	//if t.S5 != nil {
	//	size += uint64(unsafe.Sizeof(*t.S5))
	//	pointSize += 1
	//	size += uint64(len(*t.S5)) * uint64(unsafe.Sizeof(byte(0)))
	//	pointSize += 1
	//}

	return
}

func PbTestStructToBytes(uptr *uintptr, t *pb.TestStruct, pointIndex *[]uint64) ([]byte, []uint64) {
	if t == nil {
		return nil, nil
	}

	var bytes []byte
	//var pointIndex []uint64
	//pointIndex = make([]uint64, 0, 10)

	if uptr == nil {
		size, pointSize := PbTestStructToBytesSize(t)
		bytes = make([]byte, size)
		//sh := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
		//sh.Data = uintptr(jemalloc.Malloc(int(size)))
		//sh.Len = int(size)
		//sh.Cap = int(size)
		uptr = new(uintptr)
		*uptr = (*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data
		pointIndex_ := make([]uint64, 0, pointSize)
		pointIndex = &pointIndex_
		*pointIndex = append(*pointIndex, uint64(*uptr))
	}
	newT := (*pb.TestStruct)(unsafe.Pointer(*uptr))
	*uptr += unsafe.Sizeof(*t)

	//if t.IntPointField != nil {
	//	newT.IntPointField = (*int64)(unsafe.Pointer(*uptr))
	//	*newT.IntPointField = *t.IntPointField
	//	*uptr += unsafe.Sizeof(*t.IntPointField)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.IntPointField))))
	//}
	newT.IntPointField = t.IntPointField

	//if t.FloatPointField != nil {
	//	newT.FloatPointField = (*float32)(unsafe.Pointer(*uptr))
	//	*newT.FloatPointField = *t.FloatPointField
	//	*uptr += unsafe.Sizeof(*t.FloatPointField)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.FloatPointField))))
	//}
	newT.FloatPointField = t.FloatPointField

	if t.Person != nil {
		newT.Person = (*pb.Person)(unsafe.Pointer(*uptr))
		*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.Person))))
		PbPersonToBytes(uptr, t.Person, pointIndex)
	}

	//if t.IntField != nil {
	//	newT.IntField = (*int64)(unsafe.Pointer(*uptr))
	//	*newT.IntField = *t.IntField
	//	*uptr += unsafe.Sizeof(*t.IntField)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.IntField))))
	//}
	newT.IntField = t.IntField

	if len(t.IntSlice) != 0 {
		slice := (*reflect.SliceHeader)(unsafe.Pointer(&newT.IntSlice))
		slice.Cap = len(t.IntSlice)
		slice.Len = len(t.IntSlice)
		slice.Data = *uptr
		*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
		copy(newT.IntSlice, t.IntSlice)
		*uptr += uintptr(int(unsafe.Sizeof(int(0))) * len(t.IntSlice))
	}

	if len(t.Persons) != 0 {
		slice := (*reflect.SliceHeader)(unsafe.Pointer(&newT.Persons))
		slice.Cap = len(t.Persons)
		slice.Len = len(t.Persons)
		slice.Data = *uptr
		*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
		*uptr += uintptr(int(unsafe.Sizeof(int(0))) * len(t.Persons))
		for index, v := range t.Persons {
			if t.Persons != nil {
				newT.Persons[index] = (*pb.Person)(unsafe.Pointer(*uptr))
				*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.Persons[index]))))
				PbPersonToBytes(uptr, v, pointIndex)
			}
		}
	}

	//if t.I1 != nil {
	//	newT.I1 = (*int64)(unsafe.Pointer(*uptr))
	//	*newT.I1 = *t.I1
	//	*uptr += unsafe.Sizeof(*t.I1)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.I1))))
	//}
	newT.I1 = t.I1

	//if t.I2 != nil {
	//	newT.I2 = (*int64)(unsafe.Pointer(*uptr))
	//	*newT.I2 = *t.I2
	//	*uptr += unsafe.Sizeof(*t.I2)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.I2))))
	//}
	newT.I2 = t.I2

	//if t.I3 != nil {
	//	newT.I3 = (*int64)(unsafe.Pointer(*uptr))
	//	*newT.I3 = *t.I3
	//	*uptr += unsafe.Sizeof(*t.I3)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.I3))))
	//}
	newT.I3 = t.I3

	//if t.I4 != nil {
	//	newT.I4 = (*int64)(unsafe.Pointer(*uptr))
	//	*newT.I4 = *t.I4
	//	*uptr += unsafe.Sizeof(*t.I4)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.I4))))
	//}
	newT.I4 = t.I4

	//if t.I5 != nil {
	//	newT.I5 = (*int64)(unsafe.Pointer(*uptr))
	//	*newT.I5 = *t.I5
	//	*uptr += unsafe.Sizeof(*t.I5)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.I5))))
	//}
	newT.I5 = t.I5

	//if t.S1 != nil {
	//	newT.S1 = (*string)(unsafe.Pointer(*uptr))
	//	*uptr += unsafe.Sizeof(*t.S1)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.S1))))
	//	slice := (*reflect.StringHeader)(unsafe.Pointer(&*newT.S1))
	//	slice.Len = len(*t.S1)
	//	slice.Data = *uptr
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
	//	strbytes := *(*[]byte)(unsafe.Pointer(&*newT.S1))
	//	copy(strbytes, *(*[]byte)(unsafe.Pointer(&*t.S1)))
	//	*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(*t.S1))
	//}
	if len(t.S1) != 0 {
		slice := (*reflect.StringHeader)(unsafe.Pointer(&newT.S1))
		slice.Len = len(t.S1)
		slice.Data = *uptr
		*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
		strbytes := *(*[]byte)(unsafe.Pointer(&newT.S1))
		copy(strbytes, *(*[]byte)(unsafe.Pointer(&t.S1)))
		*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(t.S1))
	}

	//if t.S2 != nil {
	//	newT.S2 = (*string)(unsafe.Pointer(*uptr))
	//	*uptr += unsafe.Sizeof(*t.S2)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.S2))))
	//	slice := (*reflect.StringHeader)(unsafe.Pointer(&*newT.S2))
	//	slice.Len = len(*t.S2)
	//	slice.Data = *uptr
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
	//	strbytes := *(*[]byte)(unsafe.Pointer(&*newT.S2))
	//	copy(strbytes, *(*[]byte)(unsafe.Pointer(&*t.S2)))
	//	*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(*t.S2))
	//}
	if len(t.S2) != 0 {
		slice := (*reflect.StringHeader)(unsafe.Pointer(&newT.S2))
		slice.Len = len(t.S2)
		slice.Data = *uptr
		*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
		strbytes := *(*[]byte)(unsafe.Pointer(&newT.S2))
		copy(strbytes, *(*[]byte)(unsafe.Pointer(&t.S2)))
		*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(t.S2))
	}

	//if t.S3 != nil {
	//	newT.S3 = (*string)(unsafe.Pointer(*uptr))
	//	*uptr += unsafe.Sizeof(*t.S3)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.S3))))
	//	slice := (*reflect.StringHeader)(unsafe.Pointer(&*newT.S3))
	//	slice.Len = len(*t.S3)
	//	slice.Data = *uptr
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
	//	strbytes := *(*[]byte)(unsafe.Pointer(&*newT.S3))
	//	copy(strbytes, *(*[]byte)(unsafe.Pointer(&*t.S3)))
	//	*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(*t.S3))
	//}
	if len(t.S3) != 0 {
		slice := (*reflect.StringHeader)(unsafe.Pointer(&newT.S3))
		slice.Len = len(t.S3)
		slice.Data = *uptr
		*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
		strbytes := *(*[]byte)(unsafe.Pointer(&newT.S3))
		copy(strbytes, *(*[]byte)(unsafe.Pointer(&t.S3)))
		*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(t.S3))
	}

	//if t.S4 != nil {
	//	newT.S4 = (*string)(unsafe.Pointer(*uptr))
	//	*uptr += unsafe.Sizeof(*t.S4)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.S4))))
	//	slice := (*reflect.StringHeader)(unsafe.Pointer(&*newT.S4))
	//	slice.Len = len(*t.S4)
	//	slice.Data = *uptr
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
	//	strbytes := *(*[]byte)(unsafe.Pointer(&*newT.S4))
	//	copy(strbytes, *(*[]byte)(unsafe.Pointer(&*t.S4)))
	//	*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(*t.S4))
	//}
	if len(t.S4) != 0 {
		slice := (*reflect.StringHeader)(unsafe.Pointer(&newT.S4))
		slice.Len = len(t.S4)
		slice.Data = *uptr
		*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
		strbytes := *(*[]byte)(unsafe.Pointer(&newT.S4))
		copy(strbytes, *(*[]byte)(unsafe.Pointer(&t.S4)))
		*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(t.S4))
	}

	//if t.S5 != nil {
	//	newT.S5 = (*string)(unsafe.Pointer(*uptr))
	//	*uptr += unsafe.Sizeof(*t.S5)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newT.S5))))
	//	slice := (*reflect.StringHeader)(unsafe.Pointer(&*newT.S5))
	//	slice.Len = len(*t.S5)
	//	slice.Data = *uptr
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
	//	strbytes := *(*[]byte)(unsafe.Pointer(&*newT.S5))
	//	copy(strbytes, *(*[]byte)(unsafe.Pointer(&*t.S5)))
	//	*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(*t.S5))
	//}
	if len(t.S5) != 0 {
		slice := (*reflect.StringHeader)(unsafe.Pointer(&newT.S5))
		slice.Len = len(t.S5)
		slice.Data = *uptr
		*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
		strbytes := *(*[]byte)(unsafe.Pointer(&newT.S5))
		copy(strbytes, *(*[]byte)(unsafe.Pointer(&t.S5)))
		*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(t.S5))
	}

	return bytes, *pointIndex
}

func PbPersonToBytesSize(p *pb.Person) (size uint64, pointSize uint64) {
	if p == nil {
		return
	}

	size += uint64(unsafe.Sizeof(*p))
	pointSize += 1

	if len(p.Name) != 0 {
		size += uint64(len(p.Name)) * uint64(unsafe.Sizeof(byte(0)))
		pointSize += 1
	}

	//if p.Name != nil {
	//	size += uint64(unsafe.Sizeof(*p.Name))
	//	pointSize += 1
	//	size += uint64(len(*p.Name)) * uint64(unsafe.Sizeof(byte(0)))
	//	pointSize += 1
	//}
	//
	//if p.Age != nil {
	//	size += uint64(unsafe.Sizeof(*p.Age))
	//	pointSize += 1
	//}

	return
}

func PbPersonToBytes(uptr *uintptr, p *pb.Person, pointIndex *[]uint64) ([]byte, []uint64) {
	if p == nil {
		return nil, nil
	}

	var bytes []byte
	//var pointIndex []uint64
	//pointIndex = make([]uint64, 0, 10)

	if uptr == nil {
		size, pointSize := PbPersonToBytesSize(p)
		bytes = make([]byte, size)
		*uptr = (*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data
		pointIndex_ := make([]uint64, 0, pointSize)
		pointIndex = &pointIndex_
		*pointIndex = append(*pointIndex, uint64(*uptr))
	}

	newP := (*pb.Person)(unsafe.Pointer(*uptr))
	*uptr += unsafe.Sizeof(*newP)

	//if p.Name != nil {
	//	newP.Name = (*string)(unsafe.Pointer(*uptr))
	//	*uptr += unsafe.Sizeof(*p.Name)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newP.Name))))
	//	slice := (*reflect.StringHeader)(unsafe.Pointer(&*newP.Name))
	//	slice.Len = len(*p.Name)
	//	slice.Data = *uptr
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
	//	strbytes := *(*[]byte)(unsafe.Pointer(&*newP.Name))
	//	copy(strbytes, *(*[]byte)(unsafe.Pointer(&*p.Name)))
	//	*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(*p.Name))
	//}

	if len(p.Name) != 0 {
		slice := (*reflect.StringHeader)(unsafe.Pointer(&newP.Name))
		slice.Len = len(p.Name)
		slice.Data = *uptr
		*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&slice.Data))))
		strbytes := *(*[]byte)(unsafe.Pointer(&newP.Name))
		copy(strbytes, *(*[]byte)(unsafe.Pointer(&p.Name)))
		*uptr += uintptr(int(unsafe.Sizeof(byte(0))) * len(p.Name))
	}

	//if p.Age != nil {
	//	newP.Age = (*int64)(unsafe.Pointer(*uptr))
	//	*newP.Age = *p.Age
	//	*uptr += unsafe.Sizeof(*p.Age)
	//	*pointIndex = append(*pointIndex, uint64(uintptr(unsafe.Pointer(&newP.Age))))
	//}
	newP.Age = p.Age

	return bytes, *pointIndex
}

func BenchmarkUnMarshalWithUnsafe(b *testing.B) {
	ts := NewPbTestStruct()
	ts = NewPbTestStruct()

	bytes, pointIndex := PbTestStructToBytes(nil, ts, nil)
	cloneBytes := make([]byte, len(bytes))
	copy(cloneBytes, bytes)
	if len(pointIndex) > 1 {
		startAddr := pointIndex[0]
		cloneStartAddr := uint64((*reflect.SliceHeader)(unsafe.Pointer(&cloneBytes)).Data)
		for index, addr := range pointIndex {
			if index == 0 {
				continue
			}
			delta := byteOrder.Uint64(unsafe.Slice(&bytes[int(addr-startAddr)], 8)) - startAddr
			byteOrder.PutUint64(unsafe.Slice(&cloneBytes[int(addr-startAddr)], 8), cloneStartAddr+delta)
		}
	}
	for i := 0; i < len(bytes); i++ {
		bytes[i] = 0
	}

	newTs := (*pb.TestStruct)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&cloneBytes)).Data))
	fmt.Println(newTs)
	fmt.Println(len(bytes), len(pointIndex))

	if !proto.Equal(ts, newTs) {
		b.Fatalf("ts != newTs")
	}

	bytes, pointIndex = PbTestStructToBytes(nil, ts, nil)
	for i := 0; i < b.N; i++ {
		cloneBytes = make([]byte, len(bytes))
		copy(cloneBytes, bytes)
		if len(pointIndex) > 1 {
			startAddr := pointIndex[0]
			cloneStartAddr := uint64((*reflect.SliceHeader)(unsafe.Pointer(&cloneBytes)).Data)
			for index, addr := range pointIndex {
				if index == 0 {
					continue
				}
				delta := byteOrder.Uint64(bytes[int(addr-startAddr):int(addr-startAddr)+8]) - startAddr
				byteOrder.PutUint64(cloneBytes[int(addr-startAddr):int(addr-startAddr)+8], cloneStartAddr+delta)
			}
		}
		newTs := (*pb.TestStruct)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&cloneBytes)).Data))
		newTs.IntField = 2
	}
}

func BenchmarkMarshalWithUnsafe(b *testing.B) {
	ts := NewPbTestStruct()
	ts = NewPbTestStruct()

	//fmt.Println(ToBytesSize(ts))
	fmt.Println(PbTestStructToBytesSize(ts))

	bytes, pointIndex := PbTestStructToBytes(nil, ts, nil)
	fmt.Println("!@#", len(bytes), len(pointIndex))
	cloneBytes1 := make([]byte, len(bytes))
	copy(cloneBytes1, bytes)
	cloneBytes := make([]byte, len(cloneBytes1))
	copy(cloneBytes, cloneBytes1)
	if len(pointIndex) > 1 {
		startAddr := pointIndex[0]
		cloneStartAddr := uint64((*reflect.SliceHeader)(unsafe.Pointer(&cloneBytes)).Data)
		for index, addr := range pointIndex {
			if index == 0 {
				continue
			}
			delt := byteOrder.Uint64(bytes[int(addr-startAddr):int(addr-startAddr)+8]) - startAddr
			byteOrder.PutUint64(cloneBytes[int(addr-startAddr):int(addr-startAddr)+8], cloneStartAddr+delt)
		}
	}
	for i := 0; i < len(bytes); i++ {
		bytes[i] = 0
	}

	newTs := (*pb.TestStruct)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&cloneBytes)).Data))
	fmt.Println(newTs)
	fmt.Println(len(bytes))

	if !proto.Equal(ts, newTs) {
		b.Fatalf("ts != newTs")
	}

	for i := 0; i < b.N; i++ {
		bytes, pointIndex = PbTestStructToBytes(nil, ts, nil)
	}
}

func BenchmarkMarshalAndUnmarshalWithUnsafe(b *testing.B) {
	ts := NewPbTestStruct()
	ts = NewPbTestStruct()

	//fmt.Println(ToBytesSize(ts))
	fmt.Println(PbTestStructToBytesSize(ts))

	bytes, pointIndex := PbTestStructToBytes(nil, ts, nil)
	fmt.Println("!@#", len(bytes), len(pointIndex))
	cloneBytes := make([]byte, len(bytes))
	copy(cloneBytes, bytes)
	if len(pointIndex) > 1 {
		startAddr := pointIndex[0]
		cloneStartAddr := uint64((*reflect.SliceHeader)(unsafe.Pointer(&cloneBytes)).Data)
		for index, addr := range pointIndex {
			if index == 0 {
				continue
			}
			delt := byteOrder.Uint64(bytes[int(addr-startAddr):int(addr-startAddr)+8]) - startAddr
			byteOrder.PutUint64(cloneBytes[int(addr-startAddr):int(addr-startAddr)+8], cloneStartAddr+delt)
		}
	}
	for i := 0; i < len(bytes); i++ {
		bytes[i] = 0
	}

	newTs := (*pb.TestStruct)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&cloneBytes)).Data))
	fmt.Println(newTs)
	fmt.Println(len(bytes))

	if !proto.Equal(ts, newTs) {
		b.Fatalf("ts != newTs")
	}

	for i := 0; i < b.N; i++ {
		bytes, pointIndex := PbTestStructToBytes(nil, ts, nil)

		cloneBytes := make([]byte, len(bytes))
		copy(cloneBytes, bytes)
		if len(pointIndex) > 1 {
			startAddr := pointIndex[0]
			cloneStartAddr := uint64((*reflect.SliceHeader)(unsafe.Pointer(&cloneBytes)).Data)
			for index, addr := range pointIndex {
				if index == 0 {
					continue
				}
				delt := byteOrder.Uint64(bytes[int(addr-startAddr):int(addr-startAddr)+8]) - startAddr
				byteOrder.PutUint64(cloneBytes[int(addr-startAddr):int(addr-startAddr)+8], cloneStartAddr+delt)
			}
		}
		newTs := (*pb.TestStruct)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&cloneBytes)).Data))
		newTs.IntField = 1
	}
}

func Benchmark_make_slice(b *testing.B) {

	bytes1 := make([]byte, 1024)
	bytes2 := make([]byte, 1024)
	//bytes1[0] = 1

	for i := 0; i < b.N; i++ {
		//for j := 0; j < len(bytes1); j++ {
		//	bytes2[j] = bytes1[j]
		//}
		copy(bytes2, bytes1)
		//bytes1 = make([]byte, 1024)
	}
}
