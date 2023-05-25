package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/json-iterator/go"
	"math/rand"
	"practise/test/test_unmarshal_and_unsafe/pb"
	"testing"
	"time"
)

func NewPbTestStruct() *pb.TestStruct {
	return &pb.TestStruct{
		IntPointField:   proto.Int64(1),
		FloatPointField: proto.Float32(1.984),
		Person: &pb.Person{
			Name: proto.String("35325"),
			Age:  proto.Int64(56),
		},
		IntField: proto.Int64(34),
		IntSlice: []int64{4, 5234, 44},
		Persons: []*pb.Person{
			{
				Name: proto.String("57864"),
				Age:  proto.Int64(35),
			},
			{
				Name: proto.String("6344563"),
				Age:  proto.Int64(44),
			},
			{
				Name: proto.String("sdfgv346"),
				Age:  proto.Int64(55),
			},
			{
				Name: proto.String("45745fsds34"),
				Age:  proto.Int64(23),
			},
		},
		I1: proto.Int64(1),
		I2: proto.Int64(2),
		I3: proto.Int64(3),
		I4: proto.Int64(4),
		I5: proto.Int64(5),
		S1: proto.String("23das 2354j3wqr werfw  dasd 4523da904rh daj9023q4d "),
		S2: proto.String("5345da wer qw236fasdf a34 daiou da das098da das908"),
		S3: proto.String("73jokjdas dasdf2 adsrdas dafqwerqw342 da90 da908da"),
		S4: proto.String("6345sdfgf dassddadas sdfwrqw dasdeqdasdasgsd da908d "),
		S5: proto.String("46346dasfwefqwes asdda da jkfsfdff daj0 dahijh23rt9"),
	}
}

//func NewPbTestStruct() *pb.TestStruct {
//	return &pb.TestStruct{
//		IntPointField:   1,
//		FloatPointField: 1.984,
//		Person: &pb.Person{
//			Name: "35325",
//			Age:  56,
//		},
//		IntField: 34,
//		IntSlice: []int64{4, 5234, 44},
//		Persons: []*pb.Person{
//			{
//				Name: "57864",
//				Age:  35,
//			},
//			{
//				Name: "6344563",
//				Age:  44,
//			},
//			{
//				Name: "sdfgv346",
//				Age:  55,
//			},
//			{
//				Name: "45745fsds34",
//				Age:  23,
//			},
//		},
//		I1: 1,
//		I2: 2,
//		I3: 3,
//		I4: 4,
//		I5: 5,
//		S1: "23das 2354j3wqr werfw  dasd 4523da904rh daj9023q4d ",
//		S2: "5345da wer qw236fasdf a34 daiou da das098da das908",
//		S3: "73jokjdas dasdf2 adsrdas dafqwerqw342 da90 da908da",
//		S4: "6345sdfgf dassddadas sdfwrqw dasdeqdasdasgsd da908d ",
//		S5: "46346dasfwefqwes asdda da jkfsfdff daj0 dahijh23rt9",
//	}
//}

//func Test_abc(t *testing.T) {
//	ts := NewPbTestStruct()
//
//	v, err := proto.Marshal(ts)
//	if err != nil {
//		t.Fatalf("Marshal failed, err: %v", err)
//	}
//
//	ts = &pb.TestStruct{}
//	err = proto.Unmarshal(v, ts)
//	if err != nil {
//		t.Fatalf("Unmarshal failed, err: %v", err)
//	}
//
//	fmt.Println(ts)
//
//	bytes, _ := PbTestStructToBytes(nil, ts)
//	newTs := (*pb.TestStruct)(unsafe.Pointer((*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data))
//	fmt.Println(ts)
//	fmt.Println(newTs)
//}

func BenchmarkUnMarshal(b *testing.B) {
	ts := NewPbTestStruct()
	v, err := proto.Marshal(ts)
	if err != nil {
		b.Fatalf("Marshal failed, err: %v", err)
	}
	newTs := &pb.TestStruct{}
	proto.Unmarshal(v, newTs)
	fmt.Println(newTs)
	fmt.Println(len(v))
	for i := 0; i < b.N; i++ {
		clone := make([]byte, len(v))
		copy(clone, v)
		newTs := &pb.TestStruct{}
		proto.Unmarshal(clone, newTs)
	}
}

func BenchmarkMarshal(b *testing.B) {
	ts := NewPbTestStruct()
	v, err := proto.Marshal(ts)
	if err != nil {
		b.Fatalf("Marshal failed, err: %v", err)
	}
	newTs := &pb.TestStruct{}
	newTs.Unmarshal(v)
	fmt.Println(newTs)
	fmt.Println(len(v))
	for i := 0; i < b.N; i++ {
		//clone := make([]byte, len(v))
		//copy(clone, v)
		v, err = proto.Marshal(ts)
	}
}

func BenchmarkUnMarshal2(b *testing.B) {
	ts := NewPbTestStruct()
	v, err := ts.Marshal()
	if err != nil {
		b.Fatalf("Marshal failed, err: %v", err)
	}
	newTs := &pb.TestStruct{}
	newTs.Unmarshal(v)
	fmt.Println(newTs)
	fmt.Println(len(v))
	for i := 0; i < b.N; i++ {
		clone := make([]byte, len(v))
		copy(clone, v)
		newTs := &pb.TestStruct{}
		newTs.Unmarshal(v)
	}
}

func BenchmarkMarshal2(b *testing.B) {
	ts := NewPbTestStruct()
	v, err := ts.Marshal()
	if err != nil {
		b.Fatalf("Marshal failed, err: %v", err)
	}
	newTs := &pb.TestStruct{}
	newTs.Unmarshal(v)
	fmt.Println(newTs)
	fmt.Println(len(v))
	for i := 0; i < b.N; i++ {
		//clone := make([]byte, len(v))
		//copy(clone, v)
		v, _ = ts.Marshal()
	}
}

func BenchmarkJsonUnmarshal(b *testing.B) {
	ts := NewPbTestStruct()
	v, err := jsoniter.Marshal(ts)
	if err != nil {
		b.Fatalf("Marshal failed, err: %v", err)
	}
	newTs := &pb.TestStruct{}
	err = jsoniter.Unmarshal(v, newTs)
	if err != nil {
		b.Fatalf("Unmarshal failed, err: %v", err)
	}
	fmt.Println(newTs)
	fmt.Println(len(v))
	for i := 0; i < b.N; i++ {
		clone := make([]byte, len(v))
		copy(clone, v)
		newTs := &pb.TestStruct{}
		jsoniter.Unmarshal(v, newTs)
	}
}

func Benchmark_GogoProtobuf_Marshal(b *testing.B) {
	a := &pb.GogoProtoBufA{
		Name:     "1234567890123456",
		BirthDay: time.Now().UnixNano(),
		Phone:    "1234567890",
		Siblings: rand.Int31n(5),
		Spouse:   rand.Intn(2) == 1,
		Money:    rand.Float64(),
	}

	v, err := a.Marshal()
	if err != nil {
		b.Fatalf("Marshal failed, err: %v", err)
	}

	newA := &pb.GogoProtoBufA{}
	err = newA.Unmarshal(v)
	if err != nil {
		b.Fatalf("Unmarshal failed, err: %v", err)
	}

	fmt.Println(newA)

	for i := 0; i < b.N; i++ {
		newA = &pb.GogoProtoBufA{}
		newA.Unmarshal(v)
		//proto.Unmarshal(v, newA)
	}
}
