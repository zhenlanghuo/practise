package main

import "fmt"

func main() {
	fmt.Println(maxLength([]int{2, 2, 3, 4, 3}))
}

func maxLength(arr []int) int {
	// write code here
	ans := 1
	n := len(arr)
	m := make(map[int]int)

	l, r := 0, 0
	m[arr[0]] = 1
	for r < n-1 {
		if m[arr[r+1]] == 0 {
			r++
			m[arr[r]] = 1
			ans = max(ans, r-l+1)
		} else {
			m[arr[l]] = 0
			l++
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
