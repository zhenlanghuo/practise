package main

import (
	"fmt"
	"github.com/CodisLabs/codis/pkg/utils/unsafe2"
	"github.com/spinlock/jemalloc-go"
	"log"
	"net/http"
	_ "net/http/pprof"
	"reflect"
	"unsafe"
)

type Person struct {
	Name  string
	bytes []byte
	test  *Test
	a     int
}

type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

type Test struct {
	a int
	b int
	c int
	d int
	e int
	g int
	h int
	i int
}

func newTest(a, b int) *Test {
	return &Test{
		a: a,
		b: b,
	}
}

func main() {

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	//count := 0
	//for i := 0; i < 1e5; i++ {
	//	if test1() {
	//		count++
	//		fmt.Println("!!!!!!", i)
	//	}
	//}
	//fmt.Println(count)

	//m, s := test4()
	//defer runtime.KeepAlive(m)
	//defer runtime.KeepAlive(s)

	//m := test5()
	//defer runtime.KeepAlive(m)
	//
	//fmt.Println("=============================")
	//sigs := make(chan os.Signal, 1)
	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//<-sigs
	////time.Sleep(time.Second * )
	//fmt.Println("=============================")
	//
	//for i := 0; i < 5; i++ {
	//	start := time.Now()
	//	runtime.GC()
	//	fmt.Printf("GC took %s\n", time.Since(start))
	//}

	//p, err := New[Person]()
	//if err != nil {
	//	fmt.Println("err", err)
	//	return
	//}
	//fmt.Println(p)
	//p.Name = "dadsa111"
	//fmt.Println(p)

	a := "123"
	bytes := []byte(a)
	b := string(bytes)
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&a)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&bytes)).Data)
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&b)).Data)
}

//func test1() bool {
//	pOffheap, err := offheap.New(int64(unsafe.Sizeof(Person{})))
//	if err != nil {
//		fmt.Println("err", err)
//		return false
//	}
//	defer pOffheap.Unmap()
//	p := (*Person)(unsafe.Pointer((*SliceHeader)(unsafe.Pointer(&pOffheap)).Data))
//	//p = &Person{}
//
//	randSize := 10 + rand.Intn(100)
//	randIndex := rand.Intn(randSize)
//	randNum := byte(rand.Intn(256))
//	p.bytes = make([]byte, randSize)
//	for i := 0; i < len(p.bytes); i++ {
//		p.bytes[i] = byte(rand.Intn(256))
//	}
//	p.bytes[randIndex] = randNum
//
//	//clone := make([]byte, randSize)
//	////copy(clone, p.bytes)
//	//for i := 0; i < len(p.bytes); i++ {
//	//	clone[i] = p.bytes[i]
//	//}
//	//
//	//clone1 := make([]byte, randSize)
//	//for i := 0; i < len(clone); i++ {
//	//	clone1[i] = clone[i]
//	//}
//	//
//	//clone = nil
//
//	//fmt.Println(clone)
//	//runtime.GC()
//	//runtime.GC()
//	debug.FreeOSMemory()
//
//	//for i := 0; i < len(p.bytes); i++ {
//	//	if p.bytes[i] != clone[i] {
//	//		return true
//	//	}
//	//}
//
//	if p.bytes[randIndex] != randNum {
//		//fmt.Println(clone)
//		fmt.Println("!@#", p.bytes[randIndex], randNum)
//		return true
//	}
//
//	//if clone1[randIndex] != randNum {
//	//	return true
//	//}
//
//	return false
//}

//	func test2() bool {
//		pOffheap, err := offheap.New(int64(unsafe.Sizeof(Person{})))
//		if err != nil {
//			fmt.Println("err", err)
//			return false
//		}
//		defer pOffheap.Unmap()
//		p := (*Person)(unsafe.Pointer((*SliceHeader)(unsafe.Pointer(&pOffheap)).Data))
//
//		randSize := 1024 + rand.Intn(1024)
//		bytes := make([]byte, randSize)
//		for i := 0; i < randSize; i++ {
//			bytes[i] = byte(rand.Intn(256))
//		}
//
//		p.bytes = make([]byte, randSize)
//		if (*SliceHeader)(unsafe.Pointer(&p.bytes)).Data == (*SliceHeader)(unsafe.Pointer(&bytes)).Data {
//			fmt.Println("#######")
//		}
//		for i := 0; i < len(bytes); i++ {
//			p.bytes[i] = bytes[i]
//		}
//
//		//runtime.GC()
//		//runtime.GC()
//		debug.FreeOSMemory()
//
//		if len(p.bytes) != len(bytes) {
//			return true
//		}
//		for i := 0; i < len(p.bytes); i++ {
//			if p.bytes[i] != bytes[i] {
//				return true
//			}
//		}
//		return false
//	}
//func test3() bool {
//	pOffheap, err := offheap.New(int64(unsafe.Sizeof(Person{})))
//	if err != nil {
//		fmt.Println("err", err)
//		return false
//	}
//	defer pOffheap.Unmap()
//	p := (*Person)(unsafe.Pointer((*SliceHeader)(unsafe.Pointer(&pOffheap)).Data))
//	//p = &Person{}
//
//	a := rand.Intn(1024)
//	b := rand.Intn(1024)
//
//	//t := newTest(a, b)
//	p.test = newTest(a, b)
//	fmt.Println(a, b)
//
//	//runtime.GC()
//	//runtime.GC()
//	debug.FreeOSMemory()
//
//	if p.test.a != a || p.test.b != b {
//		return true
//	}
//
//	return false
//}

func New[T any]() (r *T, err error) {
	//fmt.Println("New size", unsafe.Sizeof(*r))
	//bytes, err := offheap.New(int64(unsafe.Sizeof(*r)))
	//n := int(unsafe.Sizeof(*r))

	//bytes := *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
	//	Data: uintptr(jemalloc.Malloc(int(unsafe.Sizeof(*r)))), Len: n, Cap: n,
	//}))
	//
	////fmt.Println(bytes, len(bytes))
	//
	//if err != nil {
	//	return
	//}
	r = (*T)(jemalloc.Malloc(int(unsafe.Sizeof(*r))))

	return
}

func NewSlice[T any](len, cap int) (r []T, err error) {
	slice := (*SliceHeader)(unsafe.Pointer(&r))
	var t T
	//fmt.Println("NewSlice size", unsafe.Sizeof(t))
	//bytes, err := offheap.New(int64(unsafe.Sizeof(t)) * int64(cap))

	bytes := unsafe2.MakeOffheapSlice(int(unsafe.Sizeof(t)) * int(cap)).Buffer()

	if err != nil {
		return
	}
	slice.Data = (*SliceHeader)(unsafe.Pointer(&bytes)).Data
	slice.Len = len
	slice.Cap = cap
	return
}

func test4() (map[int]int, []*Person) {
	m := make(map[int]int)
	slice := make([]*Person, 1e7)
	for i := 0; i < 1e7; i++ {
		//if i%1e3 == 0 {
		//	fmt.Println(i)
		//}
		p, err := New[Person]()
		if err != nil {
			fmt.Println("err", err)
			return nil, nil
		}
		//p.bytes, err = NewSlice[byte](10000, 10000)
		//p.test, err = New[Test]()
		//if err != nil {
		//	fmt.Println("err", err)
		//	return nil
		//}
		//p.bytes = make([]byte, 1)
		//p.Name = "aaa"
		p.a = i
		m[i] = i
		slice[i] = p
	}
	return m, slice
}

func test5() map[int]*Person {
	m := make(map[int]*Person)
	for i := 0; i < 1e7; i++ {
		//if i%1e3 == 0 {
		//	fmt.Println(i)
		//}
		p := &Person{}
		//p.bytes = make([]byte, 100)
		//p.test = &Test{}
		p.a = i
		m[i] = p
	}
	return m
}
