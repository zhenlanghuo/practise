package main

import "fmt"

func main() {
	fmt.Println(kmp("ababab", "abababab"))
	fmt.Println(kmp("abab", "abacabab"))
}

func kmp(S string, T string) int {
	// write code here

	n, m := len(S), len(T)
	next := make([]int, n)
	for i := 1; i < n; i++ {
		p := next[i-1]

		for p != 0 && S[p] != S[i] {
			p = next[p-1]
		}

		if S[p] == S[i] {
			p++
		}

		next[i] = p
	}

	fmt.Println(next)

	sIndex, tIndex := -1, 0
	ans := 0
	for tIndex < m {
		if S[sIndex+1] == T[tIndex] {
			sIndex++
			tIndex++
			if sIndex == n-1 {
				ans++
				sIndex = next[sIndex] - 1
			}
		} else {
			if sIndex != -1 {
				sIndex = next[sIndex] - 1
			} else {
				tIndex++
			}
		}
	}

	return ans
}
