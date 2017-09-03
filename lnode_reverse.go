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

func reverse(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}

	p := head.Next.Next
	head.Next.Next = nil
	for p != nil {
		q := p.Next
		p.Next = head.Next
		head.Next = p
		p = q
	}
}

var newList *ListNode
var endNode *ListNode

func reverseList(node *ListNode) *ListNode {
	recursiveTraverse(node)
	return newList
}

func recursiveTraverse(node *ListNode) {
	if nil == node {
		return
	}
	if nil == node.Next {
		endNode = node
		newList = endNode
		return
	}
	recursiveTraverse(node.Next)
	endNode.Next = node
	endNode = node
	endNode.Next = nil
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
		Val: 3,
	}
	cur = cur.Next
	cur.Next = &ListNode{
		Val: 4,
	}
	cur = cur.Next
	cur.Next = &ListNode{
		Val: 5,
	}
	cur = cur.Next
	cur.Next = &ListNode{
		Val: 6,
	}
	//递归调用可以没有指针
	//recursiveTraverse(n)
	//newList.Print()

	//非递归调用的需要一个头指针，val是0
	head := &ListNode{
		Next: n,
	}
	reverse(head)
	head.Print()

}
