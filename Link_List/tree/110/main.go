package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftTreeHeight := height(root.Left)
	rightTreeHeight := height(root.Right)
	return abs(leftTreeHeight-rightTreeHeight) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(height(root.Left), height(root.Right))
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		a *= -1
	}
	return a
}
