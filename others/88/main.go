package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 0, 0, 0}
	b := []int{2, 5, 6}

	fmt.Println(a, b)
	merge(a, 3, b, 3)
	fmt.Println(a)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1

	for i := m + n - 1; i >= 0; i-- {
		if index1 >= 0 && index2 >= 0 && nums1[index1] > nums2[index2] || index2 < 0 {
			nums1[i] = nums1[index1]
			index1--
		} else {
			nums1[i] = nums2[index2]
			index2--
		}
	}
}
