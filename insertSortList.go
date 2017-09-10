package main

import "fmt"
import "container/list"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	for node := l; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

func insertionSortList(head *ListNode) *ListNode {
	result := head
	result.Next = nil
	p := head.Next
	head.Print()
	for p != nil {
		if result.Next == nil {
			result.Next = p
			p = p.Next
			continue
		}
		fmt.Println(p.Val)
		p = p.Next
		// q := result.Next
		// for q != nil {
		// 	if q.Val > p.Val {
		// 		p.Next = q
		// 		q = p
		// 		break
		// 	}
		// 	q = result.Next
		// }
	}
	p.Print()
	return p
}

func main() {
	n := &ListNode{
		Val: 1,
	}
	n.Next = &ListNode{
		Val: 3,
	}
	cur := n.Next
	cur.Next = &ListNode{
		Val: 2,
	}
	cur = cur.Next
	cur.Next = &ListNode{
		Val: 5,
	}
	cur = cur.Next
	cur.Next = &ListNode{
		Val: 10,
	}
	cur = cur.Next
	cur.Next = &ListNode{
		Val: 6,
	}
	head := &ListNode{
		Next: n,
	}
	l := list.New()
	insertionSortList(n)
}
