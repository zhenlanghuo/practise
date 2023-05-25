package main

func main() {

}

func rotateMatrix(mat [][]int, n int) [][]int {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			mat[i][j], mat[j][i] = mat[j][i], mat[i][j]
		}
	}

	for i := 0; i < n; i++ {
		reverse(mat[i])
	}

	return mat
}

func reverse(array []int) {
	l, r := 0, len(array)-1
	for l < r {
		array[l], array[r] = array[r], array[l]
		l++
		r--
	}
}
