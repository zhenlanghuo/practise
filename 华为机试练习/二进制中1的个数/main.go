package main

import "fmt"

func main() {
	fmt.Println(NumberOf1(-6))
}

func NumberOf1(n int) int {
	// write code here

	n32 := int32(n)
	ans := 0

	if n32 > 0 {
		for n32 != 0 {
			if n32&1 == 1 {
				ans++
			}
			n32 = n32 >> 1
			//fmt.Println(n32)
		}
	} else {
		for n32 != 0 {
			if n32 < 0 {
				ans++
			}
			n32 = n32 << 1
			//fmt.Println(n32)
		}
	}

	return ans
}
