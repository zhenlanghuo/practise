package main

import "fmt"

func main() {
	fmt.Println(solution(81))

	fmt.Scan()
}

func solution(n int) int {
	ans := 0

	for n > 1 {
		ans += n / 3
		n = n/3 + n%3
		if n == 2 {
			n = 3
		}
	}

	return ans
}
