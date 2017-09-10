package main

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

func deleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil && current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func deleteDuplicates2(head *ListNode) *ListNode {
	current := head
	//dupli := &ListNode{}
	undupli := head
	for current != nil && current.Next != nil && current.Next.Next != nil {
		if current.Next.Next.Val == current.Next.Val {
			//fmt.Println("dupli", current.Next.Next.Val, current.Next.Val, current.Val)
			undupli = current
			fmt.Println(undupli.Val)
			current = current.Next
		} else {
			//current = current.Next
			undupli.Next = current.Next.Next
		}
		//fmt.Println("loop:", current.Next.Val, undupli.Val)
	}
	return head
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
		Val: 2,
	}
	cur := n.Next
	cur.Next = &ListNode{
		Val: 5,
	}
	cur = cur.Next
	cur.Next = &ListNode{
		Val: 5,
	}
	cur = cur.Next
	cur.Next = &ListNode{
		Val: 5,
	}
	cur = cur.Next
	cur.Next = &ListNode{
		Val: 5,
	}

	//insertionSortList(n)
	deleteDuplicates2(n)
	n.Print()
}
