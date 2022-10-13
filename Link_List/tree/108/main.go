package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helperBST(nums, 0, len(nums)-1)
}

func helperBST(nums []int, startIndex, endIndex int) *TreeNode {
	if startIndex > endIndex {
		return nil
	}
	mid := (startIndex + endIndex) / 2
	root := new(TreeNode)
	root.Val = nums[mid]
	root.Left = helperBST(nums, startIndex, mid-1)
	root.Right = helperBST(nums, mid+1, endIndex)
	return root
}
