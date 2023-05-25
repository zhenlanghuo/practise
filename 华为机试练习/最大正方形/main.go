package main

import "fmt"

func main() {
	fmt.Println(solve([][]byte{}))
}

func solve(matrix [][]byte) int {
	ans := 0

	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			ans = max(ans, dp[i][j])
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
