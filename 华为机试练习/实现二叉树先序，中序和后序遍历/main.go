package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type OrderType int

const preOrderType = 1
const midOrderType = 2
const backOrderType = 3

func threeOrders(root *TreeNode) [][]int {

	ans := make([][]int, 3)
	for i := 0; i < 3; i++ {
		ans[i] = make([]int, 0)
	}
	travel(root, preOrderType, &ans[0])
	travel(root, midOrderType, &ans[1])
	travel(root, backOrderType, &ans[2])

	return ans
}

func travel(root *TreeNode, ot OrderType, ans *[]int) {
	if root == nil {
		return
	}

	if ot == preOrderType {
		*ans = append(*ans, root.Val)
	}
	travel(root.Left, ot, ans)
	if ot == midOrderType {
		*ans = append(*ans, root.Val)
	}
	travel(root.Right, ot, ans)
	if ot == backOrderType {
		*ans = append(*ans, root.Val)
	}
}
