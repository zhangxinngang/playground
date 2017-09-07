package main

import (
	"fmt"
)

func main() {
	// a := []int{1, 2, 3, 4, 5, 6, 7}
	// b := a[5:7]
	// b = append(b, 2)
	// fmt.Println(b, a)
	r := []int{2, 0}
	merge(r, 1, []int{1}, 1)
	fmt.Println(r)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	num := m + n
	result := make([]int, num)
	idx1 := 0
	idx2 := 0
	if m == 0 {
		result = nums1
	}
	if n == 0 {
		result = nums2
	}
	if n > 0 && m > 0 {
		for i := 0; i < num; i++ {
			n1 := nums1[idx1]
			n2 := nums2[idx2]
			if n1 < n2 {
				result[i] = n2
				idx1++
			} else {
				result[i] = n1
				idx2++
			}
		}
	}
	nums1 = result
}
