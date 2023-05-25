package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, 20))
}

//type pair struct {
//	num   int
//	index int
//}

func combinationSum2(num []int, target int) [][]int {

	sort.Ints(num)
	//fmt.Println(num)

	ans := make([][]int, 0)
	var dfs func(start int, sum int, temp []int)

	dfs = func(start int, sum int, temp []int) {
		if sum > target {
			return
		}

		if sum == target {
			clone := make([]int, len(temp))
			copy(clone, temp)
			ans = append(ans, clone)
			return
		}

		if start >= len(num) {
			return
		}

		for i := start; i < len(num); i++ {
			//fmt.Println(start, sum, num[i], temp)
			dfs(i+1, sum+num[i], append(temp, num[i]))
			for i+1 < len(num) && num[i] == num[i+1] {
				i++
			}
		}
	}
	dfs(0, 0, nil)

	return ans
}
