package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)

	return result
}

func preorderTraversal2(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode
	cur := root
	for len(stack) > 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		} else {
			result = append(result, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
	}
	return result
}
