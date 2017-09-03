package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	NumMap map[int][]int
}

func Constructor(nums []int) Solution {
	nummap := map[int][]int{}
	fmt.Println(nums)
	for i, num := range nums {
		indexlist := nummap[num]
		indexlist = append(indexlist, i)
		nummap[num] = indexlist
	}
	return Solution{NumMap: nummap}
}

func (this *Solution) Pick(target int) int {
	indexList := this.NumMap[target]
	fmt.Println(this.NumMap)
	return indexList[rand.Intn(len(indexList))]
}

func main() {
	num := []int{1, 2, 3, 3, 3}
	s := Constructor(num)
	fmt.Println(s.Pick(3))
}
