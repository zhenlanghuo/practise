package main

type Student struct {
	//name    string
	age int
	//country string
	//no      string
	s1  string
	s2  string
	s3  string
	s4  string
	s5  string
	s6  string
	s7  string
	s8  string
	s9  string
	s10 string
	t1  *Student1
	t2  *Student2
}

type Student1 struct {
}

type Student2 struct {
}

//func NewStudent() *Student {
//	return &Student{
//		age: 0,
//		s1:  "1",
//		s2:  "2",
//		s3:  "3",
//		s4:  "4",
//		s5:  "5",
//		s6:  "6",
//		s7:  "7",
//		s8:  "8",
//		s9:  "9",
//		s10: "10",
//	}
//}

func main() {
	//t := test_string()
	//fmt.Println(t)
	//
	//t2 := test2()
	//fmt.Println(t2)

	//students5 := test5()

	//students6 := test6()

	//students7 := test7()

	//students8 := test8()

	//t9 := test9()

	//t10 := test10()

	//test6()

	//fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(-1))
	//
	//for i := 0; i < 1000; i++ {
	//	start := time.Now()
	//	runtime.GC()
	//	fmt.Printf("GC took %s\n", time.Since(start))
	//}

	//runtime.KeepAlive(students5)
	//runtime.KeepAlive(students6)
	//runtime.KeepAlive(students7)
	//runtime.KeepAlive(students8)
	//runtime.KeepAlive(t9)
	//runtime.KeepAlive(t10)

	//students[0].age = 1

	//test2()
	//
	//test3()

	test12()
}

//func test11() (*Student1, *Student2) {
//	return &Student1{}, &Student2{}
//}

func test12() *Student {
	return &Student{t1: &Student1{}, t2: &Student2{}, s1: "111"}
}

//func test4() []Student {
//	slice := make([]Student, 1e8)
//	for i := 0; i < 1e8; i++ {
//		slice[i] = Student{}
//	}
//	return slice
//}
//
//func test5() []*Student {
//	slice := make([]*Student, 1e7)
//	for i := 0; i < 1e7; i++ {
//		slice[i] = &Student{}
//	}
//	return slice
//}
//
//func test6() map[int]*Student {
//	m := make(map[int]*Student, 2e6)
//	for i := 0; i < 2e6; i++ {
//		m[i] = NewStudent()
//	}
//	return m
//}
//
//func test7() map[int]int {
//	m := make(map[int]int, 1e7)
//	for i := 0; i < 1e7; i++ {
//		m[i] = i
//	}
//	return m
//}
//
//func test8() map[int]Student {
//	m := make(map[int]Student, 1e7)
//	for i := 0; i < 1e7; i++ {
//		//m[i] = Student{name: "1234567890", country: "1234567890", no: "1234567890"}
//	}
//	return m
//}
//
//func test9() []string {
//	slice := make([]string, 1e8)
//	for i := 0; i < 1e8; i++ {
//		slice[i] = "1234567890"
//	}
//	return slice
//}
//
//func test10() []byte {
//	slice := make([]byte, 1e8)
//	for i := 0; i < 1e8; i++ {
//		slice = append(slice, []byte("1234567890")...)
//	}
//	return slice
//}

//func test_string() *big.Int {
//	a := new(big.Int).SetInt64(1)
//	//b := new(big.Int).SetInt64(2)
//	//c := new(big.Int).Mul(a, b)
//	return a
//}

//func test2() int64{
//	a := new(big.Int).SetInt64(1)
//	b := new(big.Int).SetInt64(2)
//	c := new(big.Int).Mul(a, b)
//	return c.Int64()
//	//return fmt.Sprintf("%v", c)
//	//fmt.Println(c)
//}
//
//func test4() string {
//	a := new(big.Int).SetInt64(1)
//	b := new(big.Int).SetInt64(2)
//	c := new(big.Int).Mul(a, b)
//	return c.String()
//	//return fmt.Sprintf("%v", c)
//	//fmt.Println(c)
//}
//
//func test3() string {
//	a := []byte("1")
//	return string(a)
//}
