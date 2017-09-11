package lnode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	for node := l; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

