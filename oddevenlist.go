package main

import (
	"fmt"
)

//Definition for singly-linked list.

//链表的所有奇数节点放在所有偶数节点的前面，奇数偶数说的是节点的位置，不是节点里面的值

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Print() {
	for node := l; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	list := []*ListNode{}
	i := 0
	for node := head; node != nil; node = node.Next {
		if i%2 != 0 {
			i++
			continue
		}
		list = append(list, node)
		i++
	}
	oddhead := list[0]
	oddcur := list[0]
	evenhead := list[0].Next
	evencur := list[0].Next
	for i := 1; i < len(list); i++ {
		node := list[i]
		oddcur.Next = node
		oddcur = node
		evencur.Next = node.Next
		evencur = node.Next
	}
	oddcur.Next = evenhead
	//evencur.Next = nil

	return oddhead
}

func main() {
	/*n := &ListNode{
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
	}*/
	list := oddEvenList(nil)
	list.Print()
}
