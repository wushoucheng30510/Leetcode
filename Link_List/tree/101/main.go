package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func isSymmetric(root *TreeNode) bool {
	return isMirrorTree(root.Left, root.Right)
}

func isMirrorTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && isMirrorTree(p.Left, q.Right) && isMirrorTree(p.Right, q.Left)
}
