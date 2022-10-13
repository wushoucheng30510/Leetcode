package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func inorderTraversal(root *TreeNode) []int {
	var output []int
	cur := root
	stack := []*TreeNode{}

	for len(stack) > 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			output = append(output, cur.Val)
			cur = cur.Right
			continue
		}
		stack = append(stack, cur)
		cur = cur.Left
	}
	return output
}

func inorderTraversal2(root *TreeNode) []int {
	var result []int

	if root == nil {
		return result
	}

	result = append(result, inorderTraversal(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorderTraversal(root.Right)...)

	return result
}
