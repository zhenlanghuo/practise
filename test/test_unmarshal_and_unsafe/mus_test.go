package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
	mus_unsafe "github.com/mus-format/mus-go/unsafe"
	"math/rand"
	"practise/test/test_unmarshal_and_unsafe/pb"
	"testing"
	"time"
)

type MUSA struct {
	Name     string
	BirthDay int64
	Phone    string
	Siblings int32
	Spouse   bool
	Money    float64
}

func MarshalMUSUnsafe(v MUSA) (buf []byte) {
	n := mus_unsafe.SizeString(v.Name)
	n += mus_unsafe.SizeInt64(v.BirthDay)
	n += mus_unsafe.SizeString(v.Phone)
	n += mus_unsafe.SizeInt32(v.Siblings)
	n += mus_unsafe.SizeBool(v.Spouse)
	n += mus_unsafe.SizeFloat64(v.Money)
	buf = make([]byte, n)
	n = mus_unsafe.MarshalString(v.Name, buf)
	n += mus_unsafe.MarshalInt64(v.BirthDay, buf[n:])
	n += mus_unsafe.MarshalString(v.Phone, buf[n:])
	n += mus_unsafe.MarshalInt32(v.Siblings, buf[n:])
	n += mus_unsafe.MarshalBool(v.Spouse, buf[n:])
	mus_unsafe.MarshalFloat64(v.Money, buf[n:])
	return
}

func UnmarshalMUSUnsafe(bs []byte) (v MUSA, n int, err error) {
	v.Name, n, err = mus_unsafe.UnmarshalString(bs)
	if err != nil {
		return
	}
	var n1 int
	v.BirthDay, n1, err = mus_unsafe.UnmarshalInt64(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Phone, n1, err = mus_unsafe.UnmarshalString(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Siblings, n1, err = mus_unsafe.UnmarshalInt32(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Spouse, n1, err = mus_unsafe.UnmarshalBool(bs[n:])
	n += n1
	if err != nil {
		return
	}
	v.Money, n1, err = mus_unsafe.UnmarshalFloat64(bs[n:])
	n += n1
	return
}

func PbTestStructMusToBytesSize(v *pb.TestStruct) (n int) {
	int64SizerFn := mus.SizerFn[int64](mus_unsafe.SizeInt64)
	//float32SizerFn := mus.SizerFn[float32](mus_unsafe.SizeFloat32)
	personSizeFn := PtrSizerFn[pb.Person](PbPersonMusToBytesSize)
	//stringSizerFn := mus.SizerFn[string](mus_unsafe.SizeString)

	//n += ord.SizePtr[int64](v.IntPointField, int64SizerFn)
	n += mus_unsafe.SizeInt64(v.IntPointField)
	//n += ord.SizePtr[float32](v.FloatPointField, float32SizerFn)
	n += mus_unsafe.SizeFloat32(v.FloatPointField)
	n += SizePtrExtern[pb.Person](v.Person, personSizeFn)
	//n += ord.SizePtr[int64](v.IntField, int64SizerFn)
	n += mus_unsafe.SizeInt64(v.IntField)
	n += ord.SizeSlice[int64](v.IntSlice, int64SizerFn)
	n += SizePtrSlice[pb.Person](v.Persons, personSizeFn)

	//n += ord.SizePtr[int64](v.I1, int64SizerFn)
	n += mus_unsafe.SizeInt64(v.I1)
	//n += ord.SizePtr[int64](v.I2, int64SizerFn)
	n += mus_unsafe.SizeInt64(v.I2)
	//n += ord.SizePtr[int64](v.I3, int64SizerFn)
	n += mus_unsafe.SizeInt64(v.I3)
	//n += ord.SizePtr[int64](v.I4, int64SizerFn)
	n += mus_unsafe.SizeInt64(v.I4)
	//n += ord.SizePtr[int64](v.I5, int64SizerFn)
	n += mus_unsafe.SizeInt64(v.I5)
	//n += ord.SizePtr[string](v.S1, stringSizerFn)
	n += mus_unsafe.SizeString(v.S1)
	//n += ord.SizePtr[string](v.S2, stringSizerFn)
	n += mus_unsafe.SizeString(v.S2)
	//n += ord.SizePtr[string](v.S3, stringSizerFn)
	n += mus_unsafe.SizeString(v.S3)
	//n += ord.SizePtr[string](v.S4, stringSizerFn)
	n += mus_unsafe.SizeString(v.S4)
	//n += ord.SizePtr[string](v.S5, stringSizerFn)
	n += mus_unsafe.SizeString(v.S5)

	return
}

func PbTestStructMusToBytes(v *pb.TestStruct, bs []byte) (n int) {
	int64MarshalerFn := mus.MarshalerFn[int64](mus_unsafe.MarshalInt64)
	//float32MarshalerFn := mus.MarshalerFn[float32](mus_unsafe.MarshalFloat32)
	personMarshalerFn := PtrMarshalerFn[pb.Person](PbPersonMusToBytes)
	//stringMarshalerFn := mus.MarshalerFn[string](mus_unsafe.MarshalString)

	//n += ord.MarshalPtr[int64](v.IntPointField, int64MarshalerFn, bs[n:])
	n += mus_unsafe.MarshalInt64(v.IntPointField, bs[n:])
	//n += ord.MarshalPtr[float32](v.FloatPointField, float32MarshalerFn, bs[n:])
	n += mus_unsafe.MarshalFloat32(v.FloatPointField, bs[n:])
	n += MarshalPtrExtern[pb.Person](v.Person, personMarshalerFn, bs[n:])
	//n += ord.MarshalPtr[int64](v.IntField, int64MarshalerFn, bs[n:])
	n += mus_unsafe.MarshalInt64(v.IntField, bs[n:])
	n += ord.MarshalSlice[int64](v.IntSlice, int64MarshalerFn, bs[n:])
	n += MarshalPtrSlice[pb.Person](v.Persons, personMarshalerFn, bs[n:])

	//n += ord.MarshalPtr[int64](v.I1, int64MarshalerFn, bs[n:])
	n += mus_unsafe.MarshalInt64(v.I1, bs[n:])
	//n += ord.MarshalPtr[int64](v.I2, int64MarshalerFn, bs[n:])
	n += mus_unsafe.MarshalInt64(v.I2, bs[n:])
	//n += ord.MarshalPtr[int64](v.I3, int64MarshalerFn, bs[n:])
	n += mus_unsafe.MarshalInt64(v.I3, bs[n:])
	//n += ord.MarshalPtr[int64](v.I4, int64MarshalerFn, bs[n:])
	n += mus_unsafe.MarshalInt64(v.I4, bs[n:])
	//n += ord.MarshalPtr[int64](v.I5, int64MarshalerFn, bs[n:])
	n += mus_unsafe.MarshalInt64(v.I5, bs[n:])

	//n += ord.MarshalPtr[string](v.S1, stringMarshalerFn, bs[n:])
	n += mus_unsafe.MarshalString(v.S1, bs[n:])
	//n += ord.MarshalPtr[string](v.S2, stringMarshalerFn, bs[n:])
	n += mus_unsafe.MarshalString(v.S2, bs[n:])
	//n += ord.MarshalPtr[string](v.S3, stringMarshalerFn, bs[n:])
	n += mus_unsafe.MarshalString(v.S3, bs[n:])
	//n += ord.MarshalPtr[string](v.S4, stringMarshalerFn, bs[n:])
	n += mus_unsafe.MarshalString(v.S4, bs[n:])
	//n += ord.MarshalPtr[string](v.S5, stringMarshalerFn, bs[n:])
	n += mus_unsafe.MarshalString(v.S5, bs[n:])

	return
}

func PbTestStructMusMarshal(v *pb.TestStruct) []byte {
	if v == nil {
		return nil
	}

	bytes := make([]byte, PbTestStructMusToBytesSize(v))
	PbTestStructMusToBytes(v, bytes)
	return bytes
}

func MusBytesToPbTestStruct(bs []byte) (v *pb.TestStruct, n int, err error) {
	int64UnmarshalerFn := mus.UnmarshalerFn[int64](mus_unsafe.UnmarshalInt64)
	//float32UnmarshalerFn := mus.UnmarshalerFn[float32](mus_unsafe.UnmarshalFloat32)
	personUnmarshalerFn := PtrUnmarshalerFn[pb.Person](MusBytesToPbPerson)
	//stringUnmarshalerFn := mus.UnmarshalerFn[string](mus_unsafe.UnmarshalString)

	v = &pb.TestStruct{}
	temp := n
	//v.IntPointField, temp, err = ord.UnmarshalPtr[int64](int64UnmarshalerFn, bs[n:])
	v.IntPointField, temp, err = mus_unsafe.UnmarshalInt64(bs[n:])
	n += temp
	//v.FloatPointField, temp, err = ord.UnmarshalPtr[float32](float32UnmarshalerFn, bs[n:])
	v.FloatPointField, temp, err = mus_unsafe.UnmarshalFloat32(bs[n:])
	n += temp
	v.Person, temp, err = UnmarshalPtrExtern[pb.Person](personUnmarshalerFn, bs[n:])
	n += temp
	//v.IntField, temp, err = ord.UnmarshalPtr[int64](int64UnmarshalerFn, bs[n:])
	v.IntField, temp, err = mus_unsafe.UnmarshalInt64(bs[n:])
	n += temp
	v.IntSlice, temp, err = ord.UnmarshalSlice[int64](int64UnmarshalerFn, bs[n:])
	n += temp
	v.Persons, temp, err = UnmarshalPtrSlice[pb.Person](personUnmarshalerFn, bs[n:])
	n += temp

	//v.I1, temp, err = ord.UnmarshalPtr[int64](int64UnmarshalerFn, bs[n:])
	v.I1, temp, err = mus_unsafe.UnmarshalInt64(bs[n:])
	n += temp
	//v.I2, temp, err = ord.UnmarshalPtr[int64](int64UnmarshalerFn, bs[n:])
	v.I2, temp, err = mus_unsafe.UnmarshalInt64(bs[n:])
	n += temp
	//v.I3, temp, err = ord.UnmarshalPtr[int64](int64UnmarshalerFn, bs[n:])
	v.I3, temp, err = mus_unsafe.UnmarshalInt64(bs[n:])
	n += temp
	//v.I4, temp, err = ord.UnmarshalPtr[int64](int64UnmarshalerFn, bs[n:])
	v.I4, temp, err = mus_unsafe.UnmarshalInt64(bs[n:])
	n += temp
	//v.I5, temp, err = ord.UnmarshalPtr[int64](int64UnmarshalerFn, bs[n:])
	v.I5, temp, err = mus_unsafe.UnmarshalInt64(bs[n:])
	n += temp

	//v.S1, temp, err = ord.UnmarshalPtr[string](stringUnmarshalerFn, bs[n:])
	v.S1, temp, err = mus_unsafe.UnmarshalString(bs[n:])
	n += temp
	//v.S2, temp, err = ord.UnmarshalPtr[string](stringUnmarshalerFn, bs[n:])
	v.S2, temp, err = mus_unsafe.UnmarshalString(bs[n:])
	n += temp
	//v.S3, temp, err = ord.UnmarshalPtr[string](stringUnmarshalerFn, bs[n:])
	v.S3, temp, err = mus_unsafe.UnmarshalString(bs[n:])
	n += temp
	//v.S4, temp, err = ord.UnmarshalPtr[string](stringUnmarshalerFn, bs[n:])
	v.S4, temp, err = mus_unsafe.UnmarshalString(bs[n:])
	n += temp
	//v.S5, temp, err = ord.UnmarshalPtr[string](stringUnmarshalerFn, bs[n:])
	v.S5, temp, err = mus_unsafe.UnmarshalString(bs[n:])
	n += temp

	return
}

func PbTestStructMusUnmarshal(bs []byte) (v *pb.TestStruct, err error) {
	ts, _, err := MusBytesToPbTestStruct(bs)
	return ts, err
}

func PbPersonMusToBytesSize(v *pb.Person) (n int) {
	//int64SizerFn := mus.SizerFn[int64](mus_unsafe.SizeInt64)
	//stringSizerFn := mus.SizerFn[string](mus_unsafe.SizeString)

	//n += ord.SizePtr[string](v.Name, stringSizerFn)
	n += mus_unsafe.SizeString(v.Name)
	//n += ord.SizePtr[int64](v.Age, int64SizerFn)
	n += mus_unsafe.SizeInt64(v.Age)

	return
}

func PbPersonMusToBytes(v *pb.Person, bs []byte) (n int) {
	//int64MarshalerFn := mus.MarshalerFn[int64](mus_unsafe.MarshalInt64)
	//stringMarshalerFn := mus.MarshalerFn[string](mus_unsafe.MarshalString)

	//n += ord.MarshalPtr[string](v.Name, stringMarshalerFn, bs[n:])
	n += mus_unsafe.MarshalString(v.Name, bs[n:])
	//n += ord.MarshalPtr[int64](v.Age, int64MarshalerFn, bs[n:])
	n += mus_unsafe.MarshalInt64(v.Age, bs[n:])

	return
}

func MusBytesToPbPerson(bs []byte) (v *pb.Person, n int, err error) {

	//int64UnmarshalerFn := mus.UnmarshalerFn[int64](mus_unsafe.UnmarshalInt64)
	//stringUnmarshalerFn := mus.UnmarshalerFn[string](mus_unsafe.UnmarshalString)

	v = &pb.Person{}
	temp := n
	//v.Name, temp, err = ord.UnmarshalPtr[string](stringUnmarshalerFn, bs[n:])
	v.Name, temp, err = mus_unsafe.UnmarshalString(bs[n:])
	n += temp
	//v.Age, temp, err = ord.UnmarshalPtr[int64](int64UnmarshalerFn, bs[n:])
	v.Age, temp, err = mus_unsafe.UnmarshalInt64(bs[n:])
	n += temp

	return
}

func Benchmark_Mus_UnMarshal(b *testing.B) {
	ts := NewPbTestStruct()

	v := PbTestStructMusMarshal(ts)
	newTs, err := PbTestStructMusUnmarshal(v)
	if err != nil {
		b.Fatalf("unmarshal failed, err: %v", err)
	}
	fmt.Println(newTs)
	fmt.Println(len(v))

	if !proto.Equal(ts, newTs) {
		b.Fatalf("ts != newTs")
	}

	for i := 0; i < b.N; i++ {
		clone := make([]byte, len(v))
		copy(clone, v)
		PbTestStructMusUnmarshal(v)
	}
}

func Benchmark_Mus_Marshal(b *testing.B) {
	ts := NewPbTestStruct()

	v := PbTestStructMusMarshal(ts)
	newTs, err := PbTestStructMusUnmarshal(v)
	if err != nil {
		b.Fatalf("unmarshal failed, err: %v", err)
	}
	fmt.Println(newTs)
	fmt.Println(len(v))

	if !proto.Equal(ts, newTs) {
		b.Fatalf("ts != newTs")
	}

	for i := 0; i < b.N; i++ {
		v = PbTestStructMusMarshal(ts)
	}
}

func Benchmark_Mus_MarshalAndUnmarshal(b *testing.B) {
	ts := NewPbTestStruct()

	v := PbTestStructMusMarshal(ts)
	newTs, err := PbTestStructMusUnmarshal(v)
	if err != nil {
		b.Fatalf("unmarshal failed, err: %v", err)
	}
	fmt.Println(newTs)
	fmt.Println(len(v))

	if !proto.Equal(ts, newTs) {
		b.Fatalf("ts != newTs")
	}

	for i := 0; i < b.N; i++ {
		v := PbTestStructMusMarshal(ts)
		PbTestStructMusUnmarshal(v)
	}
}

func Benchmark_Mus_UnMarshal_MUSA(b *testing.B) {
	musa := MUSA{
		Name:     "1234567890123456",
		BirthDay: time.Now().UnixNano(),
		Phone:    "1234567890",
		Siblings: rand.Int31n(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}

	v := MarshalMUSUnsafe(musa)
	musa, _, _ = UnmarshalMUSUnsafe(v)
	fmt.Println(musa)
	for i := 0; i < b.N; i++ {
		UnmarshalMUSUnsafe(v)
	}
}

func Benchmark_mus_unsafe(b *testing.B) {
	//a := int64(1)
	//bytes := make([]byte, mus_unsafe.SizeInt64(a))
	//mus_unsafe.MarshalInt64(a, bytes)

	s := "string4320998"
	bytes := make([]byte, mus_unsafe.SizeString(s))
	mus_unsafe.MarshalString(s, bytes)

	for i := 0; i < b.N; i++ {
		//a, _, _ = mus_unsafe.UnmarshalInt64(bytes)
		//a, _, _ = raw.UnmarshalInt64(bytes)

		//s, _, _ = mus_unsafe.UnmarshalString(bytes)
		s, _, _ = ord.UnmarshalString(bytes)
	}
}
