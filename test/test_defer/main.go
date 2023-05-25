package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	fmt.Println(f().name)

	//a := []int{}
	//year := 5
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%v, %v\n", year, math.Pow(1.035, float64(year)))
	//	year += 5
	//}
}

func f() *person {
	p := person{name: "11"}

	defer func() {
		fmt.Println(p)
		p.name = "22"
	}()

	return &p
}
