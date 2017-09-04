package main

import (
	"fmt"
)

// []int{1,2,0,4,12,0,24,0}
// []int{1,2,4,12,24,0,0,0}

func swap(list []int, i, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}

func moveZeros(nums []int) {
	lastNoneZeroFoundAt := 0
	for cur := 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			swap(nums, lastNoneZeroFoundAt, cur)
			lastNoneZeroFoundAt++
		}
		fmt.Println(nums)
	}
}

func main() {
	nums := []int{1, 0, 2, 3, 0, 4, 12, 0}
	moveZeros(nums)
	fmt.Println(nums)
}
