package main

import (
	"fmt"
)

//链表

type Node struct{
	Data int
	PrevNode *Node
	NextNode *Node
}

func main() {
	node1 := Node{
		Data:1,
	}
	node2 := Node{
		Data:2,
	}
	node1.NextNode = &node2
	node2.NextNode = &Node{
		Data:3,
	}
	for n := &node1; n != nil;n=n.NextNode{
		fmt.Println(n.Data)
	}
}