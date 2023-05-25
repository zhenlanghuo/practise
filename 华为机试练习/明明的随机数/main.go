package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)
	//fmt.Println(n)
	nums := make([]int, 0, n)
	//fmt.Println(nums)
	m := make(map[int]bool)
	temp := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&temp)
		if !m[temp] {
			nums = append(nums, temp)
		}
		m[temp] = true
	}

	sort.Ints(nums)
	for _, num := range nums {
		fmt.Println(num)
	}
}
