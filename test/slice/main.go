package main

import "fmt"

func main() {
	intSlice := []int{1, 2, 3}
	for i, v := range intSlice {
		if i == 0 {
			intSlice[1] = 3
		}
		fmt.Println(v)
	}
}

//func main() {
//	//array := make([]int, 0)
//	//array = append(array, 1)
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//array = append(array, 2)
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//array = append(array, 3)
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//array = append(array, 4)
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//array = append(array, 5)
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//array = append(array, 6)
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//array = append(array, 7)
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//
//	////array = array[1:]
//	////fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	////array = array[1:]
//	////fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	////
//	////array = append(array, 8)
//	////fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	////array = append(array, 9)
//	////fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//
//	//array = array[:len(array)-1]
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//array = array[:len(array)-1]
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//array = append(array, 8)
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//array = append(array, 9)
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//array = append(array, 10)
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//	//array = append(array, 11)
//	//fmt.Printf("array: %v, addr: %p, cap: %v, len: %v, unsafe.Pointer: %v \n", array, array, cap(array), len(array), unsafe.Pointer(&array))
//}
