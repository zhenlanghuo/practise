package main

func main() {

}

/**
 *
 * @param array int整型一维数组
 * @return int整型
 */
func FindGreatestSumOfSubArray(array []int) int {
	// write code here

	sum := 0
	ans := array[0]
	for _, num := range array {
		sum += num
		ans = max(ans, sum)
		sum = max(sum, 0)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
