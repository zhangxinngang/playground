package main

import (
	"fmt"
)

//var wg sync.WaitGroup

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 0}
	fmt.Println(isValidBST(root))
}

func removeDuplicates(nums []int) int {
	mapp := map[int]int{}
	for _, v := range nums {
		mapp[v] = 1
	}
	return len(mapp)
}

var pre = &TreeNode{}
var n = 0

func isValidBST(root *TreeNode) bool {
	if n == 0 && root.Left == nil && root.Right == nil {
		fmt.Println("sdgs")
		n++
		return true
	}
	if root != nil {
		if !isValidBST(root.Left) {
			return false
		}
		if pre != nil && root.Val <= pre.Val {
			return false
		}
		pre = root
		return isValidBST(root.Right)
	}
	return true
}
